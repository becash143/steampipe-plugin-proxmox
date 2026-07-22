package main

import (
	"github.com/becash143/steampipe-plugin-proxmox/proxmox"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: proxmox.Plugin,
	})
}
