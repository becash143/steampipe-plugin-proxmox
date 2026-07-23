# Table: proxmox_storage

Retrieve information about storage pools configured across your Proxmox VE cluster.

## Examples

### List all storage pools and their usage

```sql
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

```sql
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

### List storage pools by type

```sql
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

```sql
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

```sql
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

```sql
select
  storage,
  type,
  content
from
  proxmox_storage
where
  shared = 1;
```
