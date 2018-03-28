package rixcloud

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	BaseURL    = "https://api.rixcloud.io"
	APIVersion = "/v1/"
)

// Client is a client for making RixCloud API requests.
type Client struct {
	sling *sling.Sling
	// profile service
	Profiles *ProfileService
}

// NewClient ...
func NewClient(httpClient *http.Client, username string, password string) *Client {
	base := sling.New().Client(httpClient).Base(BaseURL + APIVersion)
	base.SetBasicAuth(username, password)
	return &Client{
		sling:    base,
		Profiles: NewProfileService(base.New()),
	}
}
