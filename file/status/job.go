package status

import (
	"context"
	"io/fs"
	"strings"
	"sync"
	"time"

	"github.com/Files-com/files-sdk-go/v2/lib"

	"github.com/hashicorp/go-retryablehttp"

	"github.com/bradfitz/iter"

	"github.com/Files-com/files-sdk-go/v2/lib/timer"

	"github.com/Files-com/files-sdk-go/v2/file/manager"
	"github.com/chilts/sid"
	ignore "github.com/sabhiram/go-gitignore"

	"github.com/Files-com/files-sdk-go/v2/directory"
	"github.com/Files-com/files-sdk-go/v2/lib/direction"

	filesSDK "github.com/Files-com/files-sdk-go/v2"
)

type EventsReporter map[Status][]Reporter

type IFile interface {
	SetStatus(Status, error)
	TransferBytes() int64
	File() filesSDK.File
	Size() int64
	Id() string
	LocalPath() string
	RemotePath() string
	Status() Status
	LastByte() time.Time
	Err() error
	Job() *Job
}

func ToStatusFile(f IFile) File {
	return File{
		TransferBytes: f.TransferBytes(),
		File:          f.File(),
		LocalPath:     f.LocalPath(),
		RemotePath:    f.RemotePath(),
		Id:            f.Id(),
		Job:           f.Job(),
		Status:        f.Status(),
		LastByte:      f.LastByte(),
		Err:           f.Err(),
	}
}

type Subscriptions struct {
	Started     chan time.Time
	Finished    chan time.Time
	Canceled    chan time.Time
	Scanning    chan time.Time
	EndScanning chan time.Time
}

type Job struct {
	Id string
	*timer.Timer
	Statuses         []IFile
	Direction        direction.Direction
	statusesMutex    *sync.RWMutex
	remoteFilesMutex *sync.Mutex
	LocalPath        string
	RemotePath       string
	Sync             bool
	CodeStart        func()
	context.CancelFunc
	cancelMutex *sync.Mutex
	Params      interface{}
	Client      interface{}
	Config      interface{}
	EventsReporter
	directory.Type
	*manager.Manager
	RetryPolicy interface{}
	*ignore.GitIgnore
	Started     *Signal
	Finished    *Signal
	Canceled    *Signal
	Scanning    *Signal
	EndScanning *Signal
	retryablehttp.Logger
	RemoteFs fs.FS
}

func (r Job) Init() *Job {
	r.statusesMutex = &sync.RWMutex{}
	r.cancelMutex = &sync.Mutex{}
	r.remoteFilesMutex = &sync.Mutex{}
	r.Id = sid.IdBase64()
	r.EventsReporter = make(map[Status][]Reporter)
	r.Timer = timer.New()
	r.Started = &Signal{}
	r.Finished = &Signal{}
	r.Canceled = &Signal{}
	r.Scanning = &Signal{}
	r.EndScanning = &Signal{}
	return &r
}

func (r *Job) SetManager(m *manager.Manager) {
	if m == nil {
		r.Manager = manager.Default()
	} else {
		r.Manager = m
	}
}

func (r *Job) SetEventsReporter(e EventsReporter) {
	if len(e) > 0 {
		r.EventsReporter = e
	}
}

func (r *Job) ClearCalled() {
	r.Started.Clear()
	r.Finished.Clear()
	r.Canceled.Clear()
	r.Scanning.Clear()
	r.EndScanning.Clear()
}

func (r *Job) ClearStatuses() Job {
	newJob := *r
	newJob.Reset()
	newJob.Statuses = []IFile{}
	return newJob
}

func (r *Job) Scan() {
	r.Scanning.call(time.Now())
}

func (r *Job) EndScan() {
	r.EndScanning.call(time.Now())
}

func (r *Job) Start(ignoreCodeStart ...bool) {
	r.Started.call(r.Timer.Start())
	if r.CodeStart != nil && len(ignoreCodeStart) == 0 {
		r.CodeStart()
	}
}

func (r *Job) Finish() {
	r.Finished.call(r.Timer.Stop())
}

func (r *Job) Cancel() {
	r.Canceled.call(r.Timer.Stop())
	r.cancelMutex.Lock()
	r.CancelFunc()
	r.cancelMutex.Unlock()
}

func (r *Job) Reset() {
	r.Timer = timer.New()
}

func (r *Job) Wait() {
	r.Finished.Wait()
}

func (r *Job) Job() *Job {
	return r
}

func (r *Job) SubscribeAll() Subscriptions {
	return Subscriptions{
		Started:     r.Started.Subscribe(),
		Finished:    r.Finished.Subscribe(),
		Canceled:    r.Canceled.Subscribe(),
		Scanning:    r.Scanning.Subscribe(),
		EndScanning: r.EndScanning.Subscribe(),
	}
}

func (r *Job) WithContext(ctx context.Context) context.Context {
	jobCtx, cancel := context.WithCancel(ctx)
	r.cancelMutex.Lock()
	r.CancelFunc = cancel
	r.cancelMutex.Unlock()
	return jobCtx
}

func (r *Job) RegisterFileEvent(callback Reporter, events ...Status) {
	for _, event := range events {
		r.EventsReporter[event] = append(r.EventsReporter[event], callback)
	}
}

type UnwrappedError struct {
	error
	OriginalError error
}

func (r *Job) UpdateStatus(status Status, file IFile, err error) {
	if err != nil && strings.Contains(err.Error(), "context canceled") {
		err = nil
		status = Canceled
	}
	file.SetStatus(status, err)
	callbacks, ok := r.EventsReporter[status]
	if ok {
		for _, callback := range callbacks {
			callback(ToStatusFile(file))
		}
	}
}

func (r *Job) Count(t ...Status) int {
	if len(t) == 0 {
		return len(r.Statuses)
	}
	var total int
	r.statusesMutex.RLock()
	for _, s := range r.Statuses {
		if s.Status().Any(t...) {
			total += 1
		}
	}
	r.statusesMutex.RUnlock()
	return total
}

func (r *Job) Add(report IFile) {
	r.statusesMutex.Lock()
	if r.EndScanning.Called() {
		panic("adding new file after Scanning is complete")
	}
	r.Statuses = append(r.Statuses, report)
	r.statusesMutex.Unlock()
}

func (r *Job) TotalBytes(t ...Status) int64 {
	var total int64
	r.statusesMutex.RLock()
	for _, s := range r.Statuses {
		if s.Status().Any(t...) {
			total += s.Size()
		}
	}
	r.statusesMutex.RUnlock()
	return total
}

func (r *Job) RemainingBytes(t ...Status) int64 {
	return r.TotalBytes(t...) - r.TransferBytes(t...)
}

func (r *Job) TransferBytes(t ...Status) int64 {
	var transferBytes int64
	r.statusesMutex.RLock()
	for _, s := range r.Statuses {
		if s.Status().Any(t...) {
			transferBytes += s.TransferBytes()
		}
	}
	r.statusesMutex.RUnlock()
	return transferBytes
}

func (r *Job) mostRecentBytes(t ...Status) (recent time.Time) {
	r.statusesMutex.RLock()
	for _, s := range r.Statuses {
		if !s.Status().Any(t...) {
			continue
		}
		if recent.IsZero() || recent.Before(s.LastByte()) {
			recent = s.LastByte()
		}
	}
	r.statusesMutex.RUnlock()
	return
}

func (r *Job) Idle(t ...Status) bool {
	return r.mostRecentBytes(t...).Before(time.Now().Add(time.Duration(-3500) * time.Millisecond))
}

func (r *Job) TransferRate(t ...Status) int64 {
	millisecondsSinceStart := time.Now().Sub(r.LastStart()).Milliseconds()
	bytesPerMilliseconds := float64(r.TransferBytes(t...)) / float64(millisecondsSinceStart)
	bytesPerSecond := bytesPerMilliseconds * float64(1000)

	if bytesPerSecond < 0 {
		return 0
	}
	return int64(bytesPerSecond)
}

func (r *Job) ETA(t ...Status) time.Duration {
	transferRate := r.TransferRate(t...)
	if transferRate == 0 {
		return 0
	}
	seconds := time.Duration(r.RemainingBytes(t...) / transferRate)
	eta := seconds * time.Second
	if eta < 0 {
		return 0
	}
	return eta
}

func (r *Job) ElapsedTime() time.Duration {
	return r.Elapsed()
}

func (r *Job) All(t ...Status) bool {
	allEnded := true
	r.statusesMutex.RLock()
	for _, s := range r.Statuses {
		if !s.Status().Any(t...) {
			allEnded = false
			break
		}
	}
	r.statusesMutex.RUnlock()
	return allEnded
}

func (r *Job) Any(t ...Status) (b bool) {
	r.statusesMutex.RLock()
	for _, s := range r.Statuses {
		if s.Status().Any(t...) {
			b = true
			break
		}
	}
	r.statusesMutex.RUnlock()
	return
}

func (r *Job) Find(t Status) (IFile, bool) {
	r.statusesMutex.RLock()
	defer r.statusesMutex.RUnlock()

	for _, s := range r.Statuses {
		if s.Status().Any(t) {
			return s, true
		}
	}

	return nil, false
}

func (r *Job) EnqueueNext() (f IFile, ok bool) {
	r.statusesMutex.Lock()
	defer func() {
		r.statusesMutex.Unlock()
		if f != nil {
			// Call UpdateStatus to run event callbacks, which needs to be done outside the mutex.
			r.UpdateStatus(Queued, f, nil)
		}
	}()

	for _, s := range r.Statuses {
		if s.Status().Any(Indexed) {
			f = s
			ok = true
			// The status must be changed within the mutex in order that it's not reused.
			s.SetStatus(Queued, nil)
			break
		}
	}

	return
}

func (r *Job) Sub(t ...Status) *Job {
	var sub []IFile
	r.statusesMutex.RLock()
	for _, s := range r.Statuses {
		if s.Status().Any(t...) {
			sub = append(sub, s)
		}
	}
	r.statusesMutex.RUnlock()
	// Causes WARNING: DATA RACE. I need to understand and fix later.
	r.cancelMutex.Lock()
	newJob := *r
	r.cancelMutex.Unlock()
	newJob.Statuses = sub
	return &newJob
}

func (r *Job) Files() []filesSDK.File {
	var files []filesSDK.File
	r.statusesMutex.RLock()
	for _, s := range r.Statuses {
		files = append(files, s.File())
	}
	r.statusesMutex.RUnlock()
	return files
}

func (r *Job) Percentage(t ...Status) int {
	p := int((float64(r.TransferBytes(t...)) / float64(r.TotalBytes(t...))) * float64(100))
	if p < 0 {
		return 0
	}
	return p
}

func (r *Job) StatusFromError(s IFile, err error) {
	if r.Canceled.Called() {
		r.UpdateStatus(Canceled, s, nil)
	} else {
		r.UpdateStatus(Errored, s, err)
	}
}

func (r *Job) FindRemoteFile(file IFile) (filesSDK.File, bool, error) {
	if r.Type == directory.File {
		entry, err := r.RemoteFs.Open(file.RemotePath())
		if err != nil {
			return filesSDK.File{}, false, err
		}
		info, err := entry.Stat()
		if err != nil {
			return filesSDK.File{}, false, err
		}

		return info.Sys().(filesSDK.File), true, nil
	} else {
		r.remoteFilesMutex.Lock()
		defer r.remoteFilesMutex.Unlock()

		dir, _ := lib.UrlLastSegment(file.RemotePath())
		entries, err := fs.ReadDir(r.RemoteFs, lib.Path{Path: dir}.ConvertEmptyToRoot().String())
		if err != nil {
			return filesSDK.File{}, false, err
		}

		for _, entry := range entries {
			if !entry.IsDir() && strings.EqualFold(lib.UrlJoinNoEscape(dir, entry.Name()), file.RemotePath()) {
				info, err := entry.Info()
				if err != nil {
					panic(err)
					return filesSDK.File{}, false, err
				}

				return info.Sys().(filesSDK.File), true, nil
			}
		}
	}

	return filesSDK.File{}, false, nil
}

func WaitTellFinished[T any](job *Job, onStatusComplete chan T, beforeCallingFinish func()) {
	event := job.EndScanning.Subscribe()
	go func() {
		wait := waitForAndCount(event, onStatusComplete)
		n := len(job.Statuses) - wait
		for range iter.N(n) {
			<-onStatusComplete
		}
		close(onStatusComplete)
		if !job.Canceled.Called() {
			beforeCallingFinish()
		}
		job.Finish()
	}()
}

func waitForAndCount[T any, F any](wait chan T, onComplete chan F) int {
	completed := 0
	for {
		select {
		case <-wait:
			return completed
		case <-onComplete:
			completed += 1
		}
	}
}
