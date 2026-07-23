# Table: proxmox_node

Retrieve information about the physical/host nodes in your Proxmox VE cluster.

## Examples

### List all nodes and their status

```sql
select
  node,
  status,
  cpu,
  maxcpu,
  mem,
  maxmem
from
  proxmox_node;
```

### Find nodes that are offline

```sql
select
  node,
  status
from
  proxmox_node
where
  status != 'online';
```

### List nodes with high memory utilization

```sql
select
  node,
  mem,
  maxmem,
  round((mem::numeric / maxmem) * 100, 2) as mem_percent_used
from
  proxmox_node
order by
  mem_percent_used desc;
```

### List nodes with high CPU utilization

```sql
select
  node,
  cpu,
  round(cpu::numeric * 100, 2) as cpu_percent_used
from
  proxmox_node
order by
  cpu_percent_used desc;
```
