---
organization: becash143
category: ["private cloud"]
icon_url: "/images/plugins/becash143/proxmox.svg"
brand_color: "#E57000"
display_name: "Proxmox VE"
name: "proxmox"
description: "Steampipe plugin for querying Proxmox VE nodes, VMs, containers, storage and more."
og_description: "Query Proxmox VE with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/becash143/proxmox-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---
# Proxmox VE + Steampipe

Proxmox VE is an open source platform for running KVM virtual machines and LXC containers. [Steampipe](https://steampipe.io) is an open source zero-ETL engine to instantly query cloud APIs using SQL.

```sql
select
  name,
  status,
  node
from
  proxmox_vm
where
  status = 'running';
```

```
+---------+---------+-------+
| name    | status  | node  |
+---------+---------+-------+
| web-01  | running | pve1  |
| db-01   | running | pve1  |
+---------+---------+-------+
```
## Documentation

- **[Table definitions & examples →](/plugins/becash143/proxmox/tables)**

## Get started

### Install

Download and install the latest Proxmox plugin:

```sh
steampipe plugin install becash143/proxmox
```

### Credentials

| Item        | Description                                                                                                   |
| ----------- | -------------------------------------------------------------------------------------------------------------- |
| Credentials | Requires an [API token](https://pve.proxmox.com/wiki/Proxmox_VE_API#API_Tokens) created in the Proxmox VE UI. |
| Permissions | The API token's associated user/role determines which resources are visible. Read-only roles are sufficient for querying. |
| Radius      | Each connection represents a single Proxmox VE cluster or standalone node.                                     |

### Configuration

Installing the latest proxmox plugin will create a config file (`~/.steampipe/config/proxmox.spc`) with a single connection named `proxmox`:

```hcl
connection "proxmox" {
  plugin = "becash143/proxmox"

  # The base URL of your Proxmox VE host, without any API path suffix.
  endpoint = "https://pve.example.com:8006"

  # API token ID, in the format "user@realm!tokenid".
  api_token = "root@pam!steampipe"

  # The secret value generated when the API token was created.
  api_secret = "1234abcd-5678-90ef-abcd-1234567890ab"

  # Set to true to skip TLS certificate verification, e.g. for self-signed certs.
  # Defaults to false.
  insecure = false
}
```

#### Creating an API token

1. Log in to the Proxmox VE web UI.
2. Navigate to **Datacenter > Permissions > API Tokens**.
3. Click **Add**, choose a user (e.g. `root@pam`), and provide a token ID (e.g. `steampipe`).
4. If you want the token scoped down instead of inheriting the user's full permissions, uncheck **Privilege Separation** only if you understand the implications; otherwise assign an appropriate role to the token via **Datacenter > Permissions**.
5. Copy the generated secret immediately — Proxmox will not show it again.

### Multiple Connections

You can query multiple Proxmox clusters or nodes at once using [aggregators](https://steampipe.io/docs/managing/connections#using-aggregators):

```hcl
connection "proxmox_dc1" {
  plugin     = "becash143/proxmox"
  endpoint   = "https://pve-dc1.example.com:8006"
  api_token  = "root@pam!steampipe"
  api_secret = "..."
}

connection "proxmox_dc2" {
  plugin     = "becash143/proxmox"
  endpoint   = "https://pve-dc2.example.com:8006"
  api_token  = "root@pam!steampipe"
  api_secret = "..."
}

connection "proxmox_all" {
  plugin      = "becash143/proxmox"
  type        = "aggregator"
  connections = ["proxmox_dc1", "proxmox_dc2"]
}
```

Querying `proxmox_all` will return results from both `proxmox_dc1` and `proxmox_dc2`.
