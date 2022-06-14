package group_user

import (
	"context"
	"strconv"

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

func (i *Iter) GroupUser() files_sdk.GroupUser {
	return i.Current().(files_sdk.GroupUser)
}

func (c *Client) List(ctx context.Context, params files_sdk.GroupUserListParams) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	path := "/group_users"
	i.ListParams = &params
	list := files_sdk.GroupUserCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list)
	return i, nil
}

func List(ctx context.Context, params files_sdk.GroupUserListParams) (*Iter, error) {
	return (&Client{}).List(ctx, params)
}

func (c *Client) Create(ctx context.Context, params files_sdk.GroupUserCreateParams) (files_sdk.GroupUser, error) {
	groupUser := files_sdk.GroupUser{}
	path := "/group_users"
	exportedParams := lib.Params{Params: params}
	data, res, err := files_sdk.Call(ctx, "POST", c.Config, path, exportedParams)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return groupUser, err
	}
	if res.StatusCode == 204 {
		return groupUser, nil
	}

	return groupUser, groupUser.UnmarshalJSON(*data)
}

func Create(ctx context.Context, params files_sdk.GroupUserCreateParams) (files_sdk.GroupUser, error) {
	return (&Client{}).Create(ctx, params)
}

func (c *Client) Update(ctx context.Context, params files_sdk.GroupUserUpdateParams) (files_sdk.GroupUser, error) {
	groupUser := files_sdk.GroupUser{}
	if params.Id == 0 {
		return groupUser, lib.CreateError(params, "Id")
	}
	path := "/group_users/" + strconv.FormatInt(params.Id, 10) + ""
	exportedParams := lib.Params{Params: params}
	data, res, err := files_sdk.Call(ctx, "PATCH", c.Config, path, exportedParams)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return groupUser, err
	}
	if res.StatusCode == 204 {
		return groupUser, nil
	}

	return groupUser, groupUser.UnmarshalJSON(*data)
}

func Update(ctx context.Context, params files_sdk.GroupUserUpdateParams) (files_sdk.GroupUser, error) {
	return (&Client{}).Update(ctx, params)
}

func (c *Client) Delete(ctx context.Context, params files_sdk.GroupUserDeleteParams) error {
	groupUser := files_sdk.GroupUser{}
	if params.Id == 0 {
		return lib.CreateError(params, "Id")
	}
	path := "/group_users/" + strconv.FormatInt(params.Id, 10) + ""
	exportedParams := lib.Params{Params: params}
	data, res, err := files_sdk.Call(ctx, "DELETE", c.Config, path, exportedParams)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return err
	}
	if res.StatusCode == 204 {
		return nil
	}

	return groupUser.UnmarshalJSON(*data)
}

func Delete(ctx context.Context, params files_sdk.GroupUserDeleteParams) error {
	return (&Client{}).Delete(ctx, params)
}
