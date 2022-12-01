package usage_daily_snapshot

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

func (i *Iter) UsageDailySnapshot() files_sdk.UsageDailySnapshot {
	return i.Current().(files_sdk.UsageDailySnapshot)
}

func (c *Client) List(ctx context.Context, params files_sdk.UsageDailySnapshotListParams, opts ...files_sdk.RequestResponseOption) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	path, err := lib.BuildPath("/usage_daily_snapshots", params)
	if err != nil {
		return i, err
	}
	i.ListParams = &params
	list := files_sdk.UsageDailySnapshotCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list, opts...)
	return i, nil
}

func List(ctx context.Context, params files_sdk.UsageDailySnapshotListParams, opts ...files_sdk.RequestResponseOption) (*Iter, error) {
	return (&Client{}).List(ctx, params, opts...)
}
