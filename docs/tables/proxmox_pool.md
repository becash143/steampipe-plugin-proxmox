# Table: proxmox_pool

Retrieve information about resource pools used to organize virtual machines, containers, and storage in Proxmox VE.

## Examples

### List all resource pools

```sql
select
  poolid,
  comment
from
  proxmox_pool;
```

### Find pools with a specific comment/description

```sql
select
  poolid,
  comment
from
  proxmox_pool
where
  comment ilike '%production%';
```

### Count members per pool

```sql
select
  poolid,
  jsonb_array_length(members) as member_count
from
  proxmox_pool;
```
