package notification

import (
	"context"
	"strconv"

	files_sdk "github.com/Files-com/files-sdk-go"
	lib "github.com/Files-com/files-sdk-go/lib"
	listquery "github.com/Files-com/files-sdk-go/listquery"
)

type Client struct {
	files_sdk.Config
}

type Iter struct {
	*lib.Iter
}

func (i *Iter) Notification() files_sdk.Notification {
	return i.Current().(files_sdk.Notification)
}

func (c *Client) List(ctx context.Context, params files_sdk.NotificationListParams) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	path := "/notifications"
	i.ListParams = &params
	list := files_sdk.NotificationCollection{}
	i.Query = listquery.Build(ctx, i, c.Config, path, &list)
	return i, nil
}

func List(ctx context.Context, params files_sdk.NotificationListParams) (*Iter, error) {
	return (&Client{}).List(ctx, params)
}

func (c *Client) Find(ctx context.Context, params files_sdk.NotificationFindParams) (files_sdk.Notification, error) {
	notification := files_sdk.Notification{}
	if params.Id == 0 {
		return notification, lib.CreateError(params, "Id")
	}
	path := "/notifications/" + strconv.FormatInt(params.Id, 10) + ""
	exportedParams, err := lib.ExportParams(params)
	if err != nil {
		return notification, err
	}
	data, res, err := files_sdk.Call(ctx, "GET", c.Config, path, exportedParams)
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return notification, err
	}
	if res.StatusCode == 204 {
		return notification, nil
	}
	if err := notification.UnmarshalJSON(*data); err != nil {
		return notification, err
	}

	return notification, nil
}

func Find(ctx context.Context, params files_sdk.NotificationFindParams) (files_sdk.Notification, error) {
	return (&Client{}).Find(ctx, params)
}

func (c *Client) Create(ctx context.Context, params files_sdk.NotificationCreateParams) (files_sdk.Notification, error) {
	notification := files_sdk.Notification{}
	path := "/notifications"
	exportedParams, err := lib.ExportParams(params)
	if err != nil {
		return notification, err
	}
	data, res, err := files_sdk.Call(ctx, "POST", c.Config, path, exportedParams)
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return notification, err
	}
	if res.StatusCode == 204 {
		return notification, nil
	}
	if err := notification.UnmarshalJSON(*data); err != nil {
		return notification, err
	}

	return notification, nil
}

func Create(ctx context.Context, params files_sdk.NotificationCreateParams) (files_sdk.Notification, error) {
	return (&Client{}).Create(ctx, params)
}

func (c *Client) Update(ctx context.Context, params files_sdk.NotificationUpdateParams) (files_sdk.Notification, error) {
	notification := files_sdk.Notification{}
	if params.Id == 0 {
		return notification, lib.CreateError(params, "Id")
	}
	path := "/notifications/" + strconv.FormatInt(params.Id, 10) + ""
	exportedParams, err := lib.ExportParams(params)
	if err != nil {
		return notification, err
	}
	data, res, err := files_sdk.Call(ctx, "PATCH", c.Config, path, exportedParams)
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return notification, err
	}
	if res.StatusCode == 204 {
		return notification, nil
	}
	if err := notification.UnmarshalJSON(*data); err != nil {
		return notification, err
	}

	return notification, nil
}

func Update(ctx context.Context, params files_sdk.NotificationUpdateParams) (files_sdk.Notification, error) {
	return (&Client{}).Update(ctx, params)
}

func (c *Client) Delete(ctx context.Context, params files_sdk.NotificationDeleteParams) (files_sdk.Notification, error) {
	notification := files_sdk.Notification{}
	if params.Id == 0 {
		return notification, lib.CreateError(params, "Id")
	}
	path := "/notifications/" + strconv.FormatInt(params.Id, 10) + ""
	exportedParams, err := lib.ExportParams(params)
	if err != nil {
		return notification, err
	}
	data, res, err := files_sdk.Call(ctx, "DELETE", c.Config, path, exportedParams)
	defer func() {
		if res != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return notification, err
	}
	if res.StatusCode == 204 {
		return notification, nil
	}
	if err := notification.UnmarshalJSON(*data); err != nil {
		return notification, err
	}

	return notification, nil
}

func Delete(ctx context.Context, params files_sdk.NotificationDeleteParams) (files_sdk.Notification, error) {
	return (&Client{}).Delete(ctx, params)
}
