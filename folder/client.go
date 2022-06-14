package folder

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

func (i *Iter) Folder() files_sdk.Folder {
	return i.Current().(files_sdk.Folder)
}

func (c *Client) ListFor(ctx context.Context, params files_sdk.FolderListForParams) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	path := lib.BuildPath("/folders/", params.Path)
	i.ListParams = &params
	list := files_sdk.FolderCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list)
	return i, nil
}

func ListFor(ctx context.Context, params files_sdk.FolderListForParams) (*Iter, error) {
	return (&Client{}).ListFor(ctx, params)
}

func (c *Client) Create(ctx context.Context, params files_sdk.FolderCreateParams) (files_sdk.File, error) {
	file := files_sdk.File{}
	path := lib.BuildPath("/folders/", params.Path)
	exportedParams := lib.Params{Params: params}
	data, res, err := files_sdk.Call(ctx, "POST", c.Config, path, exportedParams)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return file, err
	}
	if res.StatusCode == 204 {
		return file, nil
	}

	return file, file.UnmarshalJSON(*data)
}

func Create(ctx context.Context, params files_sdk.FolderCreateParams) (files_sdk.File, error) {
	return (&Client{}).Create(ctx, params)
}
