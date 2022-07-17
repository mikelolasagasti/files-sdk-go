package form_field_set

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

func (i *Iter) FormFieldSet() files_sdk.FormFieldSet {
	return i.Current().(files_sdk.FormFieldSet)
}

func (c *Client) List(ctx context.Context, params files_sdk.FormFieldSetListParams) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	path, err := lib.BuildPath("/form_field_sets", params)
	if err != nil {
		return i, err
	}
	i.ListParams = &params
	list := files_sdk.FormFieldSetCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list)
	return i, nil
}

func List(ctx context.Context, params files_sdk.FormFieldSetListParams) (*Iter, error) {
	return (&Client{}).List(ctx, params)
}

func (c *Client) Find(ctx context.Context, params files_sdk.FormFieldSetFindParams) (formFieldSet files_sdk.FormFieldSet, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "GET", Path: "/form_field_sets/{id}", Params: params, Entity: &formFieldSet})
	return
}

func Find(ctx context.Context, params files_sdk.FormFieldSetFindParams) (formFieldSet files_sdk.FormFieldSet, err error) {
	return (&Client{}).Find(ctx, params)
}

func (c *Client) Create(ctx context.Context, params files_sdk.FormFieldSetCreateParams) (formFieldSet files_sdk.FormFieldSet, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "POST", Path: "/form_field_sets", Params: params, Entity: &formFieldSet})
	return
}

func Create(ctx context.Context, params files_sdk.FormFieldSetCreateParams) (formFieldSet files_sdk.FormFieldSet, err error) {
	return (&Client{}).Create(ctx, params)
}

func (c *Client) Update(ctx context.Context, params files_sdk.FormFieldSetUpdateParams) (formFieldSet files_sdk.FormFieldSet, err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "PATCH", Path: "/form_field_sets/{id}", Params: params, Entity: &formFieldSet})
	return
}

func Update(ctx context.Context, params files_sdk.FormFieldSetUpdateParams) (formFieldSet files_sdk.FormFieldSet, err error) {
	return (&Client{}).Update(ctx, params)
}

func (c *Client) Delete(ctx context.Context, params files_sdk.FormFieldSetDeleteParams) (err error) {
	err = files_sdk.Resource(ctx, c.Config, lib.Resource{Method: "DELETE", Path: "/form_field_sets/{id}", Params: params, Entity: nil})
	return
}

func Delete(ctx context.Context, params files_sdk.FormFieldSetDeleteParams) (err error) {
	return (&Client{}).Delete(ctx, params)
}
