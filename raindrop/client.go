package raindrop

import (
	"net/http"
)

type Client struct {
	BaseURL    string
	Token      string
	HTTPClient *http.Client
	Debug      bool

	common    service
	Raindrops *RaindropsService
	Raindrop  *RaindropService
	Backup    *BackupService
	//Collections *services.CollectionService
	//Tags        *services.TagService
	//Users       *services.UserService
}

func NewClient(token string) *Client {
	c := &Client{
		BaseURL:    "https://api.raindrop.io/rest/v1",
		HTTPClient: http.DefaultClient,
		Token:      token,
		Debug:      false,
	}

	c.common.client = c
	c.Raindrops = (*RaindropsService)(&c.common)
	c.Raindrop = (*RaindropService)(&c.common)
	c.Backup = (*BackupService)(&c.common)
	//c.Collections = &services.CollectionService{Client: c}
	//c.Tags = &services.TagService{Client: c}
	//c.Users = &services.UserService{Client: c}
	return c
}
