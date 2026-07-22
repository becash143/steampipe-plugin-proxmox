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
	insecure := false

	if cfg.Insecure != nil {
		insecure = *cfg.Insecure
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecure, // #nosec G402 - user-configured TLS option
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
