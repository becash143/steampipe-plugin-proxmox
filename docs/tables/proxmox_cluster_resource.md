# Table: proxmox_cluster_resource

Retrieve a unified view of all resources (VMs, containers, storage, nodes) across the Proxmox VE cluster, as returned by the cluster resources API.

## Examples

### List all cluster resources

```sql
select
  id,
  type,
  node,
  status
from
  proxmox_cluster_resource;
```

### Filter resources by type

```sql
select
  id,
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
  status not in ('running', 'online');
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
