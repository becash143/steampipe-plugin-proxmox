---
organization: becash143
category: ["software development"]
icon_url: "/images/plugins/becash143/proxmox.svg"
brand_color: "#E57000"
display_name: Proxmox
name: proxmox
description: Steampipe plugin for querying virtual machines, containers, nodes, storage and more from Proxmox VE.
og_description: Query Proxmox VE with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/becash143/proxmox-social-graphic.png"
---

# Proxmox + Steampipe

[Proxmox VE](https://www.proxmox.com/en/proxmox-virtual-environment/overview) is an open-source server virtualization platform for managing virtual machines, containers, storage, and networking on a single solution.

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

For example:
```sql
select
  name,
  status,
  node,
  cpus,
  maxmem
from
  proxmox_vm;
```

```
+----------+---------+-------+------+------------+
| name     | status  | node  | cpus | maxmem     |
+----------+---------+-------+------+------------+
| web-01   | running | pve1  | 2    | 4294967296 |
| db-01    | stopped | pve1  | 4    | 8589934592 |
+----------+---------+-------+------+------------+
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
| Resolution  | Credentials are resolved in the following order: named profile in the connection config, environment variables. |

### Configuration

Installing the latest proxmox plugin will create a config file (`~/.steampipe/config/proxmox.spc`) with a single connection named `proxmox`:

```hcl
connection "proxmox" {
  plugin = "becash143/proxmox"

  # The base URL of your Proxmox VE API endpoint, including the /api2/json suffix.
  # Can also be set with the PROXMOX_ENDPOINT environment variable.
  endpoint = "https://pve.example.com:8006/api2/json"

  # API token ID, in the format "user@realm!tokenid".
  # Can also be set with the PROXMOX_API_TOKEN environment variable.
  api_token = "root@pam!steampipe"

  # The secret value generated when the API token was created.
  # Can also be set with the PROXMOX_API_SECRET environment variable.
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

#### Credentials from environment variables

You can also set credentials using environment variables instead of the `.spc` file:

```sh
export PROXMOX_ENDPOINT="https://pve.example.com:8006/api2/json"
export PROXMOX_API_TOKEN="root@pam!steampipe"
export PROXMOX_API_SECRET="1234abcd-5678-90ef-abcd-1234567890ab"
```

### Multiple Connections

You can query multiple Proxmox clusters or nodes at once using [aggregators](https://steampipe.io/docs/managing/connections#using-aggregators):

```hcl
connection "proxmox_dc1" {
  plugin    = "becash143/proxmox"
  endpoint  = "https://pve-dc1.example.com:8006/api2/json"
  api_token = "root@pam!steampipe"
  api_secret = "..."
}

connection "proxmox_dc2" {
  plugin    = "becash143/proxmox"
  endpoint  = "https://pve-dc2.example.com:8006/api2/json"
  api_token = "root@pam!steampipe"
  api_secret = "..."
}

connection "proxmox_all" {
  plugin      = "becash143/proxmox"
  type        = "aggregator"
  connections = ["proxmox_dc1", "proxmox_dc2"]
}
```

Querying `proxmox_all` will return results from both `proxmox_dc1` and `proxmox_dc2`.
