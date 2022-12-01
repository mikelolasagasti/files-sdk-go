package history_export_result

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

func (i *Iter) HistoryExportResult() files_sdk.HistoryExportResult {
	return i.Current().(files_sdk.HistoryExportResult)
}

func (c *Client) List(ctx context.Context, params files_sdk.HistoryExportResultListParams, opts ...files_sdk.RequestResponseOption) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	path, err := lib.BuildPath("/history_export_results", params)
	if err != nil {
		return i, err
	}
	i.ListParams = &params
	list := files_sdk.HistoryExportResultCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list, opts...)
	return i, nil
}

func List(ctx context.Context, params files_sdk.HistoryExportResultListParams, opts ...files_sdk.RequestResponseOption) (*Iter, error) {
	return (&Client{}).List(ctx, params, opts...)
}
