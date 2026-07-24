---
title: "Steampipe Table: proxmox_node - Query Proxmox VE Nodes using SQL"
description: "Allows users to query physical/host nodes in a Proxmox VE cluster for status, CPU utilization, and memory utilization."
folder: "Node"
---

# Table: proxmox_node - Query Proxmox VE Nodes using SQL

Retrieve information about the physical/host nodes in your Proxmox VE cluster.

## Table Usage Guide

The `proxmox_node` table in Steampipe provides you with information about each physical or host node within a Proxmox VE cluster. This table allows you, as a systems administrator, to query node-specific details, including online/offline status, CPU utilization, and memory utilization. You can utilize this table to gather insights on cluster health, such as identifying nodes that are offline or under heavy CPU or memory pressure.

## Examples

### List all nodes and their status

```sql+postgres
select
  node,
  status,
  cpu,
  max_cpu,
  mem,
  max_mem
from
  proxmox_node;
```

```sql+sqlite
select
  node,
  status,
  cpu,
  max_cpu,
  mem,
  max_mem
from
  proxmox_node;
```

### Find nodes that are offline

```sql+postgres
select
  node,
  status
from
  proxmox_node
where
  status != 'online';
```

```sql+sqlite
select
  node,
  status
from
  proxmox_node
where
  status != 'online';
```

### List nodes with high memory utilization

```sql+postgres
select
  node,
  mem,
  max_mem,
  round((mem::numeric / max_mem) * 100, 2) as mem_percent_used
from
  proxmox_node
order by
  mem_percent_used desc;
```

```sql+sqlite
select
  node,
  mem,
  max_mem,
  round((cast(mem as real) / max_mem) * 100, 2) as mem_percent_used
from
  proxmox_node
order by
  mem_percent_used desc;
```

### List nodes with high CPU utilization

```sql+postgres
select
  node,
  cpu,
  round(cpu::numeric * 100, 2) as cpu_percent_used
from
  proxmox_node
order by
  cpu_percent_used desc;
```

```sql+sqlite
select
  node,
  cpu,
  round(cpu * 100, 2) as cpu_percent_used
from
  proxmox_node
order by
  cpu_percent_used desc;
```
