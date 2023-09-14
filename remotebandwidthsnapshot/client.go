package remote_bandwidth_snapshot

import (
	files_sdk "github.com/Files-com/files-sdk-go/v3"
	lib "github.com/Files-com/files-sdk-go/v3/lib"
	listquery "github.com/Files-com/files-sdk-go/v3/listquery"
)

type Client struct {
	files_sdk.Config
}

type Iter struct {
	*files_sdk.Iter
	*Client
}

func (i *Iter) Reload(opts ...files_sdk.RequestResponseOption) files_sdk.IterI {
	return &Iter{Iter: i.Iter.Reload(opts...).(*files_sdk.Iter), Client: i.Client}
}

func (i *Iter) RemoteBandwidthSnapshot() files_sdk.RemoteBandwidthSnapshot {
	return i.Current().(files_sdk.RemoteBandwidthSnapshot)
}

func (c *Client) List(params files_sdk.RemoteBandwidthSnapshotListParams, opts ...files_sdk.RequestResponseOption) (*Iter, error) {
	i := &Iter{Iter: &files_sdk.Iter{}, Client: c}
	path, err := lib.BuildPath("/remote_bandwidth_snapshots", params)
	if err != nil {
		return i, err
	}
	i.ListParams = &params
	list := files_sdk.RemoteBandwidthSnapshotCollection{}
	i.Query = listquery.Build(c.Config, path, &list, opts...)
	return i, nil
}

func List(params files_sdk.RemoteBandwidthSnapshotListParams, opts ...files_sdk.RequestResponseOption) (*Iter, error) {
	return (&Client{}).List(params, opts...)
}
