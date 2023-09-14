package files_sdk

import (
	"encoding/json"

	lib "github.com/Files-com/files-sdk-go/v3/lib"
)

type App struct {
	Name                string                 `json:"name,omitempty" path:"name,omitempty" url:"name,omitempty"`
	ExtendedDescription string                 `json:"extended_description,omitempty" path:"extended_description,omitempty" url:"extended_description,omitempty"`
	ShortDescription    string                 `json:"short_description,omitempty" path:"short_description,omitempty" url:"short_description,omitempty"`
	DocumentationLinks  map[string]interface{} `json:"documentation_links,omitempty" path:"documentation_links,omitempty" url:"documentation_links,omitempty"`
	IconUrl             string                 `json:"icon_url,omitempty" path:"icon_url,omitempty" url:"icon_url,omitempty"`
	LogoUrl             string                 `json:"logo_url,omitempty" path:"logo_url,omitempty" url:"logo_url,omitempty"`
	ScreenshotListUrls  []string               `json:"screenshot_list_urls,omitempty" path:"screenshot_list_urls,omitempty" url:"screenshot_list_urls,omitempty"`
	LogoThumbnailUrl    string                 `json:"logo_thumbnail_url,omitempty" path:"logo_thumbnail_url,omitempty" url:"logo_thumbnail_url,omitempty"`
	SsoStrategyType     string                 `json:"sso_strategy_type,omitempty" path:"sso_strategy_type,omitempty" url:"sso_strategy_type,omitempty"`
	RemoteServerType    string                 `json:"remote_server_type,omitempty" path:"remote_server_type,omitempty" url:"remote_server_type,omitempty"`
	FolderBehaviorType  string                 `json:"folder_behavior_type,omitempty" path:"folder_behavior_type,omitempty" url:"folder_behavior_type,omitempty"`
	ExternalHomepageUrl string                 `json:"external_homepage_url,omitempty" path:"external_homepage_url,omitempty" url:"external_homepage_url,omitempty"`
	MarketingYoutubeUrl string                 `json:"marketing_youtube_url,omitempty" path:"marketing_youtube_url,omitempty" url:"marketing_youtube_url,omitempty"`
	TutorialYoutubeUrl  string                 `json:"tutorial_youtube_url,omitempty" path:"tutorial_youtube_url,omitempty" url:"tutorial_youtube_url,omitempty"`
	AppType             string                 `json:"app_type,omitempty" path:"app_type,omitempty" url:"app_type,omitempty"`
	Featured            *bool                  `json:"featured,omitempty" path:"featured,omitempty" url:"featured,omitempty"`
}

// Identifier no path or id

type AppCollection []App

type AppListParams struct {
	SortBy       map[string]interface{} `url:"sort_by,omitempty" required:"false" json:"sort_by,omitempty" path:"sort_by"`
	Filter       App                    `url:"filter,omitempty" required:"false" json:"filter,omitempty" path:"filter"`
	FilterPrefix map[string]interface{} `url:"filter_prefix,omitempty" required:"false" json:"filter_prefix,omitempty" path:"filter_prefix"`
	ListParams
}

func (a *App) UnmarshalJSON(data []byte) error {
	type app App
	var v app
	if err := json.Unmarshal(data, &v); err != nil {
		return lib.ErrorWithOriginalResponse{}.ProcessError(data, err, map[string]interface{}{})
	}

	*a = App(v)
	return nil
}

func (a *AppCollection) UnmarshalJSON(data []byte) error {
	type apps AppCollection
	var v apps
	if err := json.Unmarshal(data, &v); err != nil {
		return lib.ErrorWithOriginalResponse{}.ProcessError(data, err, []map[string]interface{}{})
	}

	*a = AppCollection(v)
	return nil
}

func (a *AppCollection) ToSlice() *[]interface{} {
	ret := make([]interface{}, len(*a))
	for i, v := range *a {
		ret[i] = v
	}

	return &ret
}
