# Table: proxmox_cluster_resource

Retrieve a unified view of all resources (nodes, VMs, containers, storage, pools) across the Proxmox VE cluster, as returned by the cluster resources API.

## Examples

### List all cluster resources

```sql
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

```sql
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

```sql
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

```sql
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

```sql
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

```sql
select
  name,
  type,
  node,
  mem,
  maxmem
from
  proxmox_cluster_resource
where
  type in ('qemu', 'lxc')
order by
  mem desc
limit
  10;
```
