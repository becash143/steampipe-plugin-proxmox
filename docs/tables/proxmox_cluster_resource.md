---
title: "Steampipe Table: proxmox_cluster_resource - Query Proxmox VE Cluster Resources using SQL"
description: "Allows users to query a unified view of Proxmox VE cluster resources, including nodes, VMs, containers, storage, and pools."
folder: "Cluster"
---

# Table: proxmox_cluster_resource - Query Proxmox VE Cluster Resources using SQL

Retrieve a unified view of all resources (nodes, VMs, containers, storage, pools) across the Proxmox VE cluster, as returned by the cluster resources API.

## Table Usage Guide

The `proxmox_cluster_resource` table in Steampipe provides you with information about all resources tracked by the Proxmox VE cluster resources API in a single call. This table allows you, as a systems administrator, to query resource-specific details across nodes, VMs, containers, storage, and pools without joining multiple tables, including resource type, current status, node placement, and pool membership. You can utilize this table to gather insights on resources, such as identifying resources that aren't running, counting resources by type, or finding which VMs and containers belong to a given pool.

## Examples

### List all cluster resources

```sql+postgres
select
  id,
  type,
  node,
  name,
  status
from
  proxmox_cluster_resource;
```

```sql+sqlite
select
  id,
  type,
  node,
  name,
  status
from
  proxmox_cluster_resource;
```

### Filter resources by type

```sql+postgres
select
  id,
  name,
  node,
  status
from
  proxmox_cluster_resource
where
  type = 'qemu';
```

```sql+sqlite
select
  id,
  name,
  node,
  status
from
  proxmox_cluster_resource
where
  type = 'qemu';
```

### Find resources not in a running/online state

```sql+postgres
select
  id,
  type,
  node,
  status
from
  proxmox_cluster_resource
where
  type in ('qemu', 'lxc')
  and status not in ('running', 'online');
```

```sql+sqlite
select
  id,
  type,
  node,
  status
from
  proxmox_cluster_resource
where
  type in ('qemu', 'lxc')
  and status not in ('running', 'online');
```

### Count resources by type across the cluster

```sql+postgres
select
  type,
  count(*) as resource_count
from
  proxmox_cluster_resource
group by
  type
order by
  resource_count desc;
```

```sql+sqlite
select
  type,
  count(*) as resource_count
from
  proxmox_cluster_resource
group by
  type
order by
  resource_count desc;
```

### List resources belonging to a specific pool

```sql+postgres
select
  id,
  type,
  name,
  node
from
  proxmox_cluster_resource
where
  pool = 'production';
```

```sql+sqlite
select
  id,
  type,
  name,
  node
from
  proxmox_cluster_resource
where
  pool = 'production';
```

### Find VMs/containers using the most memory

```sql+postgres
select
  name,
  type,
  node,
  mem,
  max_menm
from
  proxmox_cluster_resource
where
  type in ('qemu', 'lxc')
order by
  mem desc
limit
  10;
```

```sql+sqlite
select
  name,
  type,
  node,
  mem,
  max_mem
from
  proxmox_cluster_resource
where
  type in ('qemu', 'lxc')
order by
  mem desc
limit
  10;
```
