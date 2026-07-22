package proxmox

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Client struct {
	Endpoint   string
	TokenID    string
	TokenValue string
	Client     *http.Client
}

func NewClient(cfg Config) *Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: cfg.Insecure,
		},
	}

	return &Client{
		Endpoint:   cfg.Endpoint,
		TokenID:    cfg.APIToken,
		TokenValue: cfg.APISecret,
		Client: &http.Client{
			Transport: tr,
			Timeout:   30 * time.Second,
		},
	}
}
