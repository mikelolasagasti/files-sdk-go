package email_feedback

import (
	files_sdk "github.com/Files-com/files-sdk-go/v2"
	lib "github.com/Files-com/files-sdk-go/v2/lib"
)

type Client struct {
	files_sdk.Config
}

func (c *Client) Create(params files_sdk.EmailFeedbackCreateParams, opts ...files_sdk.RequestResponseOption) (err error) {
	err = files_sdk.Resource(c.Config, lib.Resource{Method: "POST", Path: "/email_feedback", Params: params, Entity: nil}, opts...)
	return
}

func Create(params files_sdk.EmailFeedbackCreateParams, opts ...files_sdk.RequestResponseOption) (err error) {
	return (&Client{}).Create(params, opts...)
}
