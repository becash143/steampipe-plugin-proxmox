---
title: "Steampipe Table: proxmox_pool - Query Proxmox VE Resource Pools using SQL"
description: "Allows users to query resource pools used to organize virtual machines, containers, and storage in Proxmox VE."
folder: "Pool"
---

# Table: proxmox_pool - Query Proxmox VE Resource Pools using SQL

Retrieve information about resource pools used to organize virtual machines, containers, and storage in Proxmox VE.

## Table Usage Guide

The `proxmox_pool` table in Steampipe provides you with information about resource pools within Proxmox VE. This table allows you, as a systems administrator, to query pool-specific details, including the pool identifier and its associated comment or description. You can utilize this table to gather insights on how VMs, containers, and storage are organized, such as locating pools by description or finding pools that haven't been documented with a comment.

## Examples

### List all resource pools

```sql+postgres
select
  pool_id,
  comment
from
  proxmox_pool;
```

```sql+sqlite
select
  poolid,
  comment
from
  proxmox_pool;
```

### Find pools with a specific comment/description

```sql+postgres
select
  pool_id,
  comment
from
  proxmox_pool
where
  comment ilike '%production%';
```

```sql+sqlite
select
  pool_id,
  comment
from
  proxmox_pool
where
  lower(comment) like lower('%production%');
```

### Find pools with no comment set

```sql+postgres
select
  pool_id
from
  proxmox_pool
where
  comment is null
  or comment = '';
```

```sql+sqlite
select
  pool_id
from
  proxmox_pool
where
  comment is null
  or comment = '';
```
