# Table: proxmox_vm

Retrieve information about QEMU/KVM virtual machines across your Proxmox VE cluster.

## Examples

### List all virtual machines and their current status

```sql
select
  name,
  vmid,
  node,
  status,
  cpus,
  maxmem
from
  proxmox_vm;
```

### Find all running virtual machines

```sql
select
  name,
  node,
  status
from
  proxmox_vm
where
  status = 'running';
```

### List virtual machines with more than 4 vCPUs

```sql
select
  name,
  node,
  cpus
from
  proxmox_vm
where
  cpus > 4;
```

### Count virtual machines per node

```sql
select
  node,
  count(*) as vm_count
from
  proxmox_vm
group by
  node;
```

### Find stopped virtual machines that may be reclaimable

```sql
select
  name,
  vmid,
  node,
  maxmem
from
  proxmox_vm
where
  status = 'stopped';
```
