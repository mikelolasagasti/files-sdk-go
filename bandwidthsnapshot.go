package files_sdk

import (
	"encoding/json"
	"time"

	lib "github.com/Files-com/files-sdk-go/v2/lib"
)

type BandwidthSnapshot struct {
	Id                int64      `json:"id,omitempty"`
	BytesReceived     string     `json:"bytes_received,omitempty"`
	BytesSent         string     `json:"bytes_sent,omitempty"`
	SyncBytesReceived string     `json:"sync_bytes_received,omitempty"`
	SyncBytesSent     string     `json:"sync_bytes_sent,omitempty"`
	RequestsGet       string     `json:"requests_get,omitempty"`
	RequestsPut       string     `json:"requests_put,omitempty"`
	RequestsOther     string     `json:"requests_other,omitempty"`
	LoggedAt          *time.Time `json:"logged_at,omitempty"`
}

type BandwidthSnapshotCollection []BandwidthSnapshot

type BandwidthSnapshotListParams struct {
	Cursor     string          `url:"cursor,omitempty" required:"false" json:"cursor,omitempty"`
	PerPage    int64           `url:"per_page,omitempty" required:"false" json:"per_page,omitempty"`
	SortBy     json.RawMessage `url:"sort_by,omitempty" required:"false" json:"sort_by,omitempty"`
	Filter     json.RawMessage `url:"filter,omitempty" required:"false" json:"filter,omitempty"`
	FilterGt   json.RawMessage `url:"filter_gt,omitempty" required:"false" json:"filter_gt,omitempty"`
	FilterGteq json.RawMessage `url:"filter_gteq,omitempty" required:"false" json:"filter_gteq,omitempty"`
	FilterLike json.RawMessage `url:"filter_like,omitempty" required:"false" json:"filter_like,omitempty"`
	FilterLt   json.RawMessage `url:"filter_lt,omitempty" required:"false" json:"filter_lt,omitempty"`
	FilterLteq json.RawMessage `url:"filter_lteq,omitempty" required:"false" json:"filter_lteq,omitempty"`
	lib.ListParams
}

func (b *BandwidthSnapshot) UnmarshalJSON(data []byte) error {
	type bandwidthSnapshot BandwidthSnapshot
	var v bandwidthSnapshot
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BandwidthSnapshot(v)
	return nil
}

func (b *BandwidthSnapshotCollection) UnmarshalJSON(data []byte) error {
	type bandwidthSnapshots []BandwidthSnapshot
	var v bandwidthSnapshots
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BandwidthSnapshotCollection(v)
	return nil
}

func (b *BandwidthSnapshotCollection) ToSlice() *[]interface{} {
	ret := make([]interface{}, len(*b))
	for i, v := range *b {
		ret[i] = v
	}

	return &ret
}
