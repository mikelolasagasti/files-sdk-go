package site

import (
	"context"

	files_sdk "github.com/Files-com/files-sdk-go/v2"
	lib "github.com/Files-com/files-sdk-go/v2/lib"
)

type Client struct {
	files_sdk.Config
}

func (c *Client) Get(ctx context.Context, opts ...files_sdk.RequestResponseOption) (site files_sdk.Site, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "GET", Path: "/site", Params: lib.Interface(), Entity: &site}, opts...)
	return
}

func Get(ctx context.Context, opts ...files_sdk.RequestResponseOption) (site files_sdk.Site, err error) {
	return (&Client{}).Get(ctx, opts...)
}

func (c *Client) GetUsage(ctx context.Context, opts ...files_sdk.RequestResponseOption) (usageSnapshot files_sdk.UsageSnapshot, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "GET", Path: "/site/usage", Params: lib.Interface(), Entity: &usageSnapshot}, opts...)
	return
}

func GetUsage(ctx context.Context, opts ...files_sdk.RequestResponseOption) (usageSnapshot files_sdk.UsageSnapshot, err error) {
	return (&Client{}).GetUsage(ctx, opts...)
}

func (c *Client) Update(ctx context.Context, params files_sdk.SiteUpdateParams, opts ...files_sdk.RequestResponseOption) (site files_sdk.Site, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "PATCH", Path: "/site", Params: params, Entity: &site}, opts...)
	return
}

func Update(ctx context.Context, params files_sdk.SiteUpdateParams, opts ...files_sdk.RequestResponseOption) (site files_sdk.Site, err error) {
	return (&Client{}).Update(ctx, params, opts...)
}
