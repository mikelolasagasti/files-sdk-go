package files_sdk

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hashicorp/go-retryablehttp"
)

//go:embed _VERSION
var VERSION string

const (
	UserAgent   = "Files.com Go SDK"
	DefaultSite = "app"
	APIPath     = "/api/rest/v1"
)

var GlobalConfig Config

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
	Get(string) (*http.Response, error)
}

type Logger interface {
	Printf(string, ...interface{})
}

type Config struct {
	APIKey            string `header:"X-FilesAPI-Key"`
	SessionId         string `header:"X-FilesAPI-Auth"`
	Endpoint          string
	Subdomain         string
	standardClient    HttpClient
	rawClient         *retryablehttp.Client
	AdditionalHeaders map[string]string
	logger            Logger
	Debug             *bool
	UserAgent         string
	Environment
}

func (c *Config) SetHttpClient(client *http.Client) {
	c.GetRawClient().HTTPClient = client
}

func (c *Config) GetHttpClient() HttpClient {
	if c.standardClient == nil {
		c.standardClient = c.GetRawClient().StandardClient()
	}
	return c.standardClient
}

func (c *Config) GetRawClient() *retryablehttp.Client {
	if c.rawClient == nil {
		c.rawClient = retryablehttp.NewClient()
		c.rawClient.ErrorHandler = retryablehttp.PassthroughErrorHandler
		c.rawClient.Logger = c.Logger()
		c.rawClient.RetryMax = 3
	}

	return c.rawClient
}

type NullLogger struct{}

func (n NullLogger) Printf(_ string, _ ...interface{}) {
}

func (c *Config) InDebug() bool {
	return c.Debug != nil && *c.Debug || (os.Getenv("FILES_SDK_DEBUG") != "")
}

func (c *Config) Logger() retryablehttp.Logger {
	if c.logger != nil {
		return c.logger
	}
	if c.InDebug() {
		c.SetLogger(log.New(os.Stderr, "", log.LstdFlags))
	} else {
		c.SetLogger(NullLogger{})
	}
	return c.logger
}

func (c *Config) SetLogger(l Logger) {
	c.logger = l
}

func (c *Config) RootPath() string {
	if c.Subdomain == "" {
		c.Subdomain = DefaultSite
	}
	if c.Endpoint == "" {
		c.Endpoint = strings.Replace(c.Environment.Endpoint(), "{SUBDOMAIN}", c.Subdomain, 1)
	}
	return c.Endpoint + APIPath
}

func (c *Config) GetAPIKey() string {
	if c.APIKey != "" {
		return c.APIKey
	}
	if os.Getenv("FILES_API_KEY") != "" {
		return os.Getenv("FILES_API_KEY")
	}
	return ""
}

func (c *Config) SetHeaders(headers *http.Header) {
	if c.UserAgent == "" {
		c.UserAgent = fmt.Sprintf("%v %v", UserAgent, strings.TrimSpace(VERSION))
	}
	headers.Set("User-Agent", c.UserAgent)
	if c.GetAPIKey() != "" {
		headers.Set("X-FilesAPI-Key", c.GetAPIKey())
	} else if c.SessionId != "" {
		headers.Set("X-FilesAPI-Auth", c.SessionId)
	}

	for key, value := range c.AdditionalHeaders {
		headers.Set(key, value)
	}
}
