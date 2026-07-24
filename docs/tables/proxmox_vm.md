---
title: "Steampipe Table: proxmox_vm - Query Proxmox VE Virtual Machines using SQL"
description: "Allows users to query QEMU/KVM virtual machines across a Proxmox VE cluster, including status, CPU allocation, and memory allocation."
folder: "VM"
---

# Table: proxmox_vm - Query Proxmox VE Virtual Machines using SQL

Retrieve information about QEMU/KVM virtual machines across your Proxmox VE cluster.

## Table Usage Guide

The `proxmox_vm` table in Steampipe provides you with information about QEMU/KVM virtual machines within Proxmox VE. This table allows you, as a systems administrator, to query VM-specific details, including current status, node placement, vCPU allocation, and memory allocation. You can utilize this table to gather insights on VMs, such as identifying running versus stopped machines, finding VMs sized above a given vCPU threshold, or counting VMs per node.

## Examples

### List all virtual machines and their current status

```sql+postgres
select
  name,
  vm_id,
  node,
  status,
  cpus,
  max_mem
from
  proxmox_vm;
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
  proxmox_vm;
```

### Find all running virtual machines

```sql+postgres
select
  name,
  node,
  status
from
  proxmox_vm
where
  status = 'running';
```

```sql+sqlite
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

```sql+postgres
select
  name,
  node,
  cpus
from
  proxmox_vm
where
  cpus > 4;
```

```sql+sqlite
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

```sql+postgres
select
  node,
  count(*) as vm_count
from
  proxmox_vm
group by
  node;
```

```sql+sqlite
select
  node,
  count(*) as vm_count
from
  proxmox_vm
group by
  node;
```

### Find stopped virtual machines that may be reclaimable

```sql+postgres
select
  name,
  vm_id,
  node,
  max_mem
from
  proxmox_vm
where
  status = 'stopped';
```

```sql+sqlite
select
  name,
  vm_id,
  node,
  max_mem
from
  proxmox_vm
where
  status = 'stopped';
```
