---
title: "Steampipe Table: proxmox_storage - Query Proxmox VE Storage Pools using SQL"
description: "Allows users to query storage pools configured across a Proxmox VE cluster, including usage, type, content, and shared/active state."
folder: "Storage"
---

# Table: proxmox_storage - Query Proxmox VE Storage Pools using SQL

Retrieve information about storage pools configured across your Proxmox VE cluster.

## Table Usage Guide

The `proxmox_storage` table in Steampipe provides you with information about storage pools configured within Proxmox VE. This table allows you, as a systems administrator, to query storage-specific details, including capacity, usage, storage type, content types supported, and whether a pool is shared across nodes or currently active. You can utilize this table to gather insights on storage, such as identifying pools nearing capacity, finding inactive pools, or locating pools shared across the cluster.

## Examples

### List all storage pools and their usage

```sql+postgres
select
  storage,
  type,
  used,
  total,
  avail
from
  proxmox_storage;
```

```sql+sqlite
select
  storage,
  type,
  used,
  total,
  avail
from
  proxmox_storage;
```

### Find storage pools that are nearly full

```sql+postgres
select
  storage,
  type,
  round((used::numeric / total) * 100, 2) as percent_used
from
  proxmox_storage
where
  total > 0
  and (used::numeric / total) > 0.85
order by
  percent_used desc;
```

```sql+sqlite
select
  storage,
  type,
  round((cast(used as real) / total) * 100, 2) as percent_used
from
  proxmox_storage
where
  total > 0
  and (cast(used as real) / total) > 0.85
order by
  percent_used desc;
```

### List storage pools by type

```sql+postgres
select
  storage,
  type,
  content
from
  proxmox_storage
order by
  type;
```

```sql+sqlite
select
  storage,
  type,
  content
from
  proxmox_storage
order by
  type;
```

### Find storage pools with the most available space

```sql+postgres
select
  storage,
  type,
  avail
from
  proxmox_storage
order by
  avail desc
limit
  5;
```

```sql+sqlite
select
  storage,
  type,
  avail
from
  proxmox_storage
order by
  avail desc
limit
  5;
```

### Find inactive storage pools

```sql+postgres
select
  storage,
  type,
  active
from
  proxmox_storage
where
  active = 0;
```

```sql+sqlite
select
  storage,
  type,
  active
from
  proxmox_storage
where
  active = 0;
```

### Find storage pools shared across nodes

```sql+postgres
select
  storage,
  type,
  content
from
  proxmox_storage
where
  shared = 1;
```

```sql+sqlite
select
  storage,
  type,
  content
from
  proxmox_storage
where
  shared = 1;
```
