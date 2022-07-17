package bandwidth_snapshot

import (
	"context"

	files_sdk "github.com/Files-com/files-sdk-go/v2"
	lib "github.com/Files-com/files-sdk-go/v2/lib"
	listquery "github.com/Files-com/files-sdk-go/v2/listquery"
)

type Client struct {
	files_sdk.Config
}

type Iter struct {
	*lib.Iter
}

func (i *Iter) BandwidthSnapshot() files_sdk.BandwidthSnapshot {
	return i.Current().(files_sdk.BandwidthSnapshot)
}

func (c *Client) List(ctx context.Context, params files_sdk.BandwidthSnapshotListParams) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	path, err := lib.BuildPath("/bandwidth_snapshots", params)
	if err != nil {
		return i, err
	}
	i.ListParams = &params
	list := files_sdk.BandwidthSnapshotCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list)
	return i, nil
}

func List(ctx context.Context, params files_sdk.BandwidthSnapshotListParams) (*Iter, error) {
	return (&Client{}).List(ctx, params)
}
