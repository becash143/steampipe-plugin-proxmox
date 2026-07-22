package proxmox

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type apiResponse[T any] struct {
	Data T `json:"data"`
}

// doRequest performs an authenticated GET request to the Proxmox API.
func (c *Client) doRequest(path string, result any) error {
	req, err := http.NewRequest(http.MethodGet, c.Endpoint+path, nil)
	if err != nil {
		return err
	}

	req.Header.Set(
		"Authorization",
		fmt.Sprintf("PVEAPIToken=%s=%s", c.TokenID, c.TokenValue),
	)

	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("proxmox API returned %d: %s", resp.StatusCode, string(body))
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

// ListNodes returns all Proxmox nodes.
func (c *Client) ListNodes() ([]Node, error) {
	var response apiResponse[[]Node]

	if err := c.doRequest("/api2/json/nodes", &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
