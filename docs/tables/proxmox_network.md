---
title: "Steampipe Table: proxmox_network - Query Proxmox VE Network Interfaces using SQL"
description: "Allows users to query network interfaces configured on Proxmox VE nodes, including bridges, addressing, and boot-time configuration."
folder: "Network"
---

# Table: proxmox_network - Query Proxmox VE Network Interfaces using SQL

Retrieve information about network interfaces configured on nodes in your Proxmox VE cluster.

## Table Usage Guide

The `proxmox_network` table in Steampipe provides you with information about network interfaces configured on each node in a Proxmox VE cluster. This table allows you, as a systems administrator, to query interface-specific details, including interface type, addressing method, active/autostart state, and node placement. You can utilize this table to gather insights on networking, such as identifying inactive interfaces, auditing which interfaces rely on DHCP versus static addressing, and counting interfaces per node.

## Examples

### List all network interfaces

```sql+postgres
select
  iface,
  node,
  type,
  active,
  address
from
  proxmox_network;
```

```sql+sqlite
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

```sql+postgres
select
  iface,
  node,
  type
from
  proxmox_network
where
  active = 0;
```

```sql+sqlite
select
  iface,
  node,
  type
from
  proxmox_network
where
  active = 0;
```

### List bridge interfaces

```sql+postgres
select
  iface,
  node,
  address,
  gateway
from
  proxmox_network
where
  type = 'bridge';
```

```sql+sqlite
select
  iface,
  node,
  address,
  gateway
from
  proxmox_network
where
  type = 'bridge';
```

### Count network interfaces per node

```sql+postgres
select
  node,
  count(*) as interface_count
from
  proxmox_network
group by
  node;
```

```sql+sqlite
select
  node,
  count(*) as interface_count
from
  proxmox_network
group by
  node;
```

### Find interfaces using DHCP instead of static IPs

```sql+postgres
select
  iface,
  node,
  method
from
  proxmox_network
where
  method = 'dhcp';
```

```sql+sqlite
select
  iface,
  node,
  method
from
  proxmox_network
where
  method = 'dhcp';
```

### Find interfaces that won't start automatically on boot

```sql+postgres
select
  iface,
  node,
  type
from
  proxmox_network
where
  autostart = 0;
```

```sql+sqlite
select
  iface,
  node,
  type
from
  proxmox_network
where
  autostart = 0;
```
