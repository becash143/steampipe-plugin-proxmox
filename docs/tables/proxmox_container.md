---
title: "Steampipe Table: proxmox_container - Query Proxmox VE LXC Containers using SQL"
description: "Allows users to query Proxmox VE LXC containers for status, resource allocation, node placement, and more."
folder: "Container"
---

# Table: proxmox_container - Query Proxmox VE LXC Containers using SQL

Retrieve information about LXC containers across your Proxmox VE cluster.

## Table Usage Guide

The `proxmox_container` table in Steampipe provides you with information about LXC containers within Proxmox VE. This table allows you, as a systems administrator, to query container-specific details, including status, CPU allocation, memory allocation, and node placement. You can utilize this table to gather insights on containers, such as which containers are running or stopped, how containers are distributed across nodes, and which containers are allocated the most memory.

## Examples

### List all containers and their current status

```sql+postgres
select
  name,
  vm_id,
  node,
  status,
  cpus,
  max_mem
from
  proxmox_container;
```

```sql+sqlite
select
  name,
  vm_id,
  node,
  status,
  cpus,
  max_mem
from
  proxmox_container;
```

### Find all running containers

```sql+postgres
select
  name,
  node,
  status
from
  proxmox_container
where
  status = 'running';
```

```sql+sqlite
select
  name,
  node,
  status
from
  proxmox_container
where
  status = 'running';
```

### Count containers per node

```sql+postgres
select
  node,
  count(*) as container_count
from
  proxmox_container
group by
  node;
```

```sql+sqlite
select
  node,
  count(*) as container_count
from
  proxmox_container
group by
  node;
```

### List containers using more than 2GB of allocated memory

```sql+postgres
select
  name,
  node,
  max_mem
from
  proxmox_container
where
  max_mem > 2147483648;
```

```sql+sqlite
select
  name,
  node,
  max_mem
from
  proxmox_container
where
  max_mem > 2147483648;
```

### Find stopped containers

```sql+postgres
select
  name,
  vm_id,
  node
from
  proxmox_container
where
  status = 'stopped';
```

```sql+sqlite
select
  name,
  vm_id,
  node
from
  proxmox_container
where
  status = 'stopped';
```
