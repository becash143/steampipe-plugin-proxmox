---
title: "Steampipe Table: proxmox_user - Query Proxmox VE Users using SQL"
description: "Allows users to query Proxmox VE access-control users, including is_enabled state, expiration, realm, and contact details."
folder: "User"
---

# Table: proxmox_user - Query Proxmox VE Users using SQL

Retrieve information about users configured in your Proxmox VE cluster's access control.

## Table Usage Guide

The `proxmox_user` table in Steampipe provides you with information about users configured in Proxmox VE's access control system. This table allows you, as a systems administrator, to query user-specific details, including enabled/disabled state, account expiration, realm membership, and contact information. You can utilize this table to gather insights on access control, such as identifying disabled accounts, auditing which users have an expiration date set, or finding users missing an email address.

## Examples

### List all users

```sql+postgres
select
  user_id,
  is_enabled,
  expire,
  comment
from
  proxmox_user;
```

```sql+sqlite
select
  user_id,
  is_enabled,
  expire,
  comment
from
  proxmox_user;
```

### Find disabled user accounts

```sql+postgres
select
  user_id,
  comment
from
  proxmox_user
where
  not is_enabled;
```

```sql+sqlite
select
  user_id,
  comment
from
  proxmox_user
where
  not is_enabled;
```

### Find users with an expiration date set

```sql+postgres
select
  user_id,
  expire
from
  proxmox_user
where
  expire is not null;
```

```sql+sqlite
select
  user_id,
  expire
from
  proxmox_user
where
  expire is not null;
```

### Find users belonging to a specific realm

```sql+postgres
select
  user_id,
  comment
from
  proxmox_user
where
  user_id like '%@pam';
```

```sql+sqlite
select
  user_id,
  comment
from
  proxmox_user
where
  user_id like '%@pam';
```

### Find users missing an email address

```sql+postgres
select
  user_id,
  comment
from
  proxmox_user
where
  email is null
  or email = '';
```

```sql+sqlite
select
  user_id,
  comment
from
  proxmox_user
where
  email is null
  or email = '';
```
