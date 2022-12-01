package invoice

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

func (i *Iter) AccountLineItem() files_sdk.AccountLineItem {
	return i.Current().(files_sdk.AccountLineItem)
}

func (c *Client) List(ctx context.Context, params files_sdk.InvoiceListParams, opts ...files_sdk.RequestResponseOption) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	path, err := lib.BuildPath("/invoices", params)
	if err != nil {
		return i, err
	}
	i.ListParams = &params
	list := files_sdk.AccountLineItemCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list, opts...)
	return i, nil
}

func List(ctx context.Context, params files_sdk.InvoiceListParams, opts ...files_sdk.RequestResponseOption) (*Iter, error) {
	return (&Client{}).List(ctx, params, opts...)
}

func (c *Client) Find(ctx context.Context, params files_sdk.InvoiceFindParams, opts ...files_sdk.RequestResponseOption) (accountLineItem files_sdk.AccountLineItem, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "GET", Path: "/invoices/{id}", Params: params, Entity: &accountLineItem}, opts...)
	return
}

func Find(ctx context.Context, params files_sdk.InvoiceFindParams, opts ...files_sdk.RequestResponseOption) (accountLineItem files_sdk.AccountLineItem, err error) {
	return (&Client{}).Find(ctx, params, opts...)
}
