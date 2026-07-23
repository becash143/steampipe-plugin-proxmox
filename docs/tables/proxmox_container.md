# Table: proxmox_container

Retrieve information about LXC containers across your Proxmox VE cluster.

## Examples

### List all containers and their current status

```sql
select
  name,
  vmid,
  node,
  status,
  cpus,
  maxmem
from
  proxmox_container;
```

### Find all running containers

```sql
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

```sql
select
  node,
  count(*) as container_count
from
  proxmox_container
group by
  node;
```

### List containers using more than 2GB of allocated memory

```sql
select
  name,
  node,
  maxmem
from
  proxmox_container
where
  maxmem > 2147483648;
```

### Find stopped containers

```sql
select
  name,
  vmid,
  node
from
  proxmox_container
where
  status = 'stopped';
```
