# Steampipe Plugin: Proxmox

Use SQL to query your Proxmox VE cluster — nodes, virtual machines, LXC containers, storage, networking, tasks, users, and resource pools.

[Steampipe](https://steampipe.io) is an open source CLI that exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. This plugin adds Proxmox VE as a queryable data source.

## Installation

### Option 1 — Install from a published release (recommended)

```bash
steampipe plugin install ghcr.io/becash143/proxmox
```

### Option 2 — Build and install locally from source

```bash
git clone https://github.com/becash143/steampipe-plugin-proxmox.git
cd steampipe-plugin-proxmox
go build -o steampipe-plugin-proxmox.plugin .
mkdir -p ~/.steampipe/plugins/local/proxmox
cp steampipe-plugin-proxmox.plugin ~/.steampipe/plugins/local/proxmox/steampipe-plugin-proxmox.plugin
chmod +x ~/.steampipe/plugins/local/proxmox/steampipe-plugin-proxmox.plugin
```

When installed locally, reference the plugin as `local/proxmox` in your connection config (see below).

## Configuration

Create `~/.steampipe/config/proxmox.spc` with your connection details.

### Authenticate with an API token (recommended)

```hcl
connection "proxmox" {
  plugin     = "local/proxmox"          # or "ghcr.io/becash143/proxmox" if installed via Option 1
  endpoint   = "https://your-proxmox-host:8006"
  api_token  = "user@realm!tokenid"     # e.g. "bikash@pve!inventory"
  api_secret = "your-token-secret"
  insecure   = true                     # set false if using a trusted/valid TLS certificate
}
```

Create an API token in the Proxmox web UI under **Datacenter → Permissions → API Tokens**. For a read-only inventory tool like this plugin, assign the token the **PVEAuditor** role at path `/`, and make sure **Privilege Separation** is unchecked (or an explicit role is granted) so the token can see VMs, containers, and other resources — not just nodes.



| Argument     | Type   | Required | Description                                                                 |
|--------------|--------|----------|-------------------------------------------------------------------------------|
| `endpoint`   | string | Yes      | Base URL of your Proxmox API, e.g. `https://proxmox.example.com:8006`         |
| `api_token`  | string | No*      | API token ID in `user@realm!tokenid` format                                   |
| `api_secret` | string | No*      | API token secret                                                               |
| `insecure`   | bool   | No       | Skip TLS certificate verification. Defaults to `false`.                        |

\*  The token pair is required.

## Get Involved

* Source code: [github.com/becash143/steampipe-plugin-proxmox](https://github.com/becash143/steampipe-plugin-proxmox)
* Steampipe: [steampipe.io](https://steampipe.io)
* Issues and feature requests: open a GitHub issue on this repo

## Tables

| Table                        | Description                                                        |
|-------------------------------|---------------------------------------------------------------------|
| `proxmox_node`                 | Physical nodes in the Proxmox cluster                                |
| `proxmox_vm`                    | QEMU virtual machines across all nodes                                |
| `proxmox_container`             | LXC containers across all nodes                                        |
| `proxmox_cluster_resource`      | Unified view of nodes, VMs, containers, storage, and pools in one call |
| `proxmox_storage`               | Configured storage pools and their usage                                |
| `proxmox_network`               | Network interfaces on each node                                         |
| `proxmox_task`                  | Recent task/job history per node                                         |
| `proxmox_user`                  | Proxmox access-control users                                              |
| `proxmox_pool`                  | Resource pools                                                              |

## Example Queries

**All running VMs, sorted by memory usage:**
```sql
select
  name,
  node,
  status,
  mem,
  maxmem
from
  proxmox_vm
where
  status = 'running'
order by
  mem desc;
```

**VM and container count per node:**
```sql
select
  node,
  count(*) filter (where type = 'qemu') as vms,
  count(*) filter (where type = 'lxc') as containers
from
  proxmox_cluster_resource
where
  type in ('qemu', 'lxc')
group by
  node;
```

**Storage pools nearing capacity:**
```sql
select
  storage,
  type,
  used,
  total,
  round((used::numeric / nullif(total, 0)) * 100, 1) as percent_used
from
  proxmox_storage
where
  total > 0
order by
  percent_used desc;
```

**Failed tasks in the last day:**
```sql
select
  node,
  type,
  user,
  status,
  to_timestamp(start_time) as started_at
from
  proxmox_task
where
  status != 'OK'
  and start_time > extract(epoch from now() - interval '1 day')
order by
  start_time desc;
```

## Development

Requires [Go](https://go.dev) 1.26+ and the [Steampipe CLI](https://steampipe.io/downloads).

```bash
git clone https://github.com/becash143/steampipe-plugin-proxmox.git
cd steampipe-plugin-proxmox
go build -o steampipe-plugin-proxmox.plugin .
```

Copy the resulting binary to `~/.steampipe/plugins/local/proxmox/steampipe-plugin-proxmox.plugin`, restart the Steampipe service, and query away.

## License

Apache License 2.0
