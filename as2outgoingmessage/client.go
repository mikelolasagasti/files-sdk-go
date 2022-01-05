package as2_outgoing_message

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

func (i *Iter) As2OutgoingMessage() files_sdk.As2OutgoingMessage {
	return i.Current().(files_sdk.As2OutgoingMessage)
}

func (c *Client) List(ctx context.Context, params files_sdk.As2OutgoingMessageListParams) (*Iter, error) {
	i := &Iter{Iter: &lib.Iter{}}
	params.ListParams.Set(params.Page, params.PerPage, params.Cursor, params.MaxPages)
	path := "/as2_outgoing_messages"
	i.ListParams = &params
	list := files_sdk.As2OutgoingMessageCollection{}
	i.Query = listquery.Build(ctx, i, c.Config, path, &list)
	return i, nil
}

func List(ctx context.Context, params files_sdk.As2OutgoingMessageListParams) (*Iter, error) {
	return (&Client{}).List(ctx, params)
}
