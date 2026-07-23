# Table: proxmox_network

Retrieve information about network interfaces configured on nodes in your Proxmox VE cluster.

## Examples

### List all network interfaces

```sql
select
  iface,
  node,
  type,
  active,
  address
from
  proxmox_network;
```

### Find inactive network interfaces

```sql
select
  iface,
  node,
  type
from
  proxmox_network
where
  active = false;
```

### List bridge interfaces

```sql
select
  iface,
  node,
  address
from
  proxmox_network
where
  type = 'bridge';
```

### Count network interfaces per node

```sql
select
  node,
  count(*) as interface_count
from
  proxmox_network
group by
  node;
```
