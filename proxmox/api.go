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

// ListVMs returns all QEMU VMs on a given node.
func (c *Client) ListVMs(node string) ([]VM, error) {
	var response apiResponse[[]VM]
	if err := c.doRequest(fmt.Sprintf("/api2/json/nodes/%s/qemu", node), &response); err != nil {
		return nil, err
	}
	for i := range response.Data {
		response.Data[i].Node = node
	}
	return response.Data, nil
}

// ListContainers returns all LXC containers on a given node.
func (c *Client) ListContainers(node string) ([]Container, error) {
	var response apiResponse[[]Container]
	if err := c.doRequest(fmt.Sprintf("/api2/json/nodes/%s/lxc", node), &response); err != nil {
		return nil, err
	}
	for i := range response.Data {
		response.Data[i].Node = node
	}
	return response.Data, nil
}

// ListClusterResources returns a unified view of all cluster resources.
func (c *Client) ListClusterResources() ([]ClusterResource, error) {
	var response apiResponse[[]ClusterResource]
	if err := c.doRequest("/api2/json/cluster/resources", &response); err != nil {
		return nil, err
	}
	return response.Data, nil
}

// ListStorage returns all storage pools.
func (c *Client) ListStorage() ([]Storage, error) {
	var response apiResponse[[]Storage]
	if err := c.doRequest("/api2/json/storage", &response); err != nil {
		return nil, err
	}
	return response.Data, nil
}

// ListNetworkInterfaces returns network interfaces on a given node.
func (c *Client) ListNetworkInterfaces(node string) ([]NetworkInterface, error) {
	var response apiResponse[[]NetworkInterface]
	if err := c.doRequest(fmt.Sprintf("/api2/json/nodes/%s/network", node), &response); err != nil {
		return nil, err
	}
	for i := range response.Data {
		response.Data[i].Node = node
	}
	return response.Data, nil
}

// ListTasks returns recent tasks on a given node.
func (c *Client) ListTasks(node string) ([]Task, error) {
	var response apiResponse[[]Task]
	if err := c.doRequest(fmt.Sprintf("/api2/json/nodes/%s/tasks", node), &response); err != nil {
		return nil, err
	}
	for i := range response.Data {
		response.Data[i].Node = node
	}
	return response.Data, nil
}

// ListUsers returns all Proxmox users.
func (c *Client) ListUsers() ([]User, error) {
	var response apiResponse[[]User]
	if err := c.doRequest("/api2/json/access/users", &response); err != nil {
		return nil, err
	}
	return response.Data, nil
}

// ListPools returns all resource pools.
func (c *Client) ListPools() ([]Pool, error) {
	var response apiResponse[[]Pool]
	if err := c.doRequest("/api2/json/pools", &response); err != nil {
		return nil, err
	}
	return response.Data, nil
}
