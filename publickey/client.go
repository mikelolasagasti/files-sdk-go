package public_key

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

func (i *Iter) PublicKey() files_sdk.PublicKey {
	return i.Current().(files_sdk.PublicKey)
}

func (c *Client) List(ctx context.Context, params files_sdk.PublicKeyListParams) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	path := "/public_keys"
	i.ListParams = &params
	list := files_sdk.PublicKeyCollection{}
	i.Query = listquery.Build(ctx, c.Config, path, &list)
	return i, nil
}

func List(ctx context.Context, params files_sdk.PublicKeyListParams) (*Iter, error) {
	return (&Client{}).List(ctx, params)
}

func (c *Client) Find(ctx context.Context, params files_sdk.PublicKeyFindParams) (files_sdk.PublicKey, error) {
	publicKey := files_sdk.PublicKey{}
	if params.Id == 0 {
		return publicKey, lib.CreateError(params, "Id")
	}
	path := "/public_keys/" + strconv.FormatInt(params.Id, 10) + ""
	exportedParams := lib.Params{Params: params}
	data, res, err := files_sdk.Call(ctx, "GET", c.Config, path, exportedParams)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return publicKey, err
	}
	if res.StatusCode == 204 {
		return publicKey, nil
	}

	return publicKey, publicKey.UnmarshalJSON(*data)
}

func Find(ctx context.Context, params files_sdk.PublicKeyFindParams) (files_sdk.PublicKey, error) {
	return (&Client{}).Find(ctx, params)
}

func (c *Client) Create(ctx context.Context, params files_sdk.PublicKeyCreateParams) (files_sdk.PublicKey, error) {
	publicKey := files_sdk.PublicKey{}
	path := "/public_keys"
	exportedParams := lib.Params{Params: params}
	data, res, err := files_sdk.Call(ctx, "POST", c.Config, path, exportedParams)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return publicKey, err
	}
	if res.StatusCode == 204 {
		return publicKey, nil
	}

	return publicKey, publicKey.UnmarshalJSON(*data)
}

func Create(ctx context.Context, params files_sdk.PublicKeyCreateParams) (files_sdk.PublicKey, error) {
	return (&Client{}).Create(ctx, params)
}

func (c *Client) Update(ctx context.Context, params files_sdk.PublicKeyUpdateParams) (files_sdk.PublicKey, error) {
	publicKey := files_sdk.PublicKey{}
	if params.Id == 0 {
		return publicKey, lib.CreateError(params, "Id")
	}
	path := "/public_keys/" + strconv.FormatInt(params.Id, 10) + ""
	exportedParams := lib.Params{Params: params}
	data, res, err := files_sdk.Call(ctx, "PATCH", c.Config, path, exportedParams)
	defer func() {
		if res != nil && res.Body != nil {
			res.Body.Close()
		}
	}()
	if err != nil {
		return publicKey, err
	}
	if res.StatusCode == 204 {
		return publicKey, nil
	}

	return publicKey, publicKey.UnmarshalJSON(*data)
}

func Update(ctx context.Context, params files_sdk.PublicKeyUpdateParams) (files_sdk.PublicKey, error) {
	return (&Client{}).Update(ctx, params)
}

func (c *Client) Delete(ctx context.Context, params files_sdk.PublicKeyDeleteParams) error {
	publicKey := files_sdk.PublicKey{}
	if params.Id == 0 {
		return lib.CreateError(params, "Id")
	}
	path := "/public_keys/" + strconv.FormatInt(params.Id, 10) + ""
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

	return publicKey.UnmarshalJSON(*data)
}

func Delete(ctx context.Context, params files_sdk.PublicKeyDeleteParams) error {
	return (&Client{}).Delete(ctx, params)
}
