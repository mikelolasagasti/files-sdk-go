package notification

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

func (i *Iter) Notification() files_sdk.Notification {
	return i.Current().(files_sdk.Notification)
}

func (c *Client) List(ctx context.Context, params files_sdk.NotificationListParams) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	path, err := lib.BuildPath("/notifications", params)
	if err != nil {
		return i, err
	}
	i.ListParams = &params
	list := files_sdk.NotificationCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list)
	return i, nil
}

func List(ctx context.Context, params files_sdk.NotificationListParams) (*Iter, error) {
	return (&Client{}).List(ctx, params)
}

func (c *Client) Find(ctx context.Context, params files_sdk.NotificationFindParams) (notification files_sdk.Notification, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "GET", Path: "/notifications/{id}", Params: params, Entity: &notification})
	return
}

func Find(ctx context.Context, params files_sdk.NotificationFindParams) (notification files_sdk.Notification, err error) {
	return (&Client{}).Find(ctx, params)
}

func (c *Client) Create(ctx context.Context, params files_sdk.NotificationCreateParams) (notification files_sdk.Notification, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "POST", Path: "/notifications", Params: params, Entity: &notification})
	return
}

func Create(ctx context.Context, params files_sdk.NotificationCreateParams) (notification files_sdk.Notification, err error) {
	return (&Client{}).Create(ctx, params)
}

func (c *Client) Update(ctx context.Context, params files_sdk.NotificationUpdateParams) (notification files_sdk.Notification, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "PATCH", Path: "/notifications/{id}", Params: params, Entity: &notification})
	return
}

func Update(ctx context.Context, params files_sdk.NotificationUpdateParams) (notification files_sdk.Notification, err error) {
	return (&Client{}).Update(ctx, params)
}

func (c *Client) Delete(ctx context.Context, params files_sdk.NotificationDeleteParams) (err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "DELETE", Path: "/notifications/{id}", Params: params, Entity: nil})
	return
}

func Delete(ctx context.Context, params files_sdk.NotificationDeleteParams) (err error) {
	return (&Client{}).Delete(ctx, params)
}
