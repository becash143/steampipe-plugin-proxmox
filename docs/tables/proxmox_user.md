# Table: proxmox_user

Retrieve information about users configured in your Proxmox VE cluster's access control.

## Examples

### List all users

```sql
select
  userid,
  enable,
  expire,
  comment
from
  proxmox_user;
```

### Find disabled user accounts

```sql
select
  userid,
  comment
from
  proxmox_user
where
  enable = false;
```

### Find users with an expiration date set

```sql
select
  userid,
  expire
from
  proxmox_user
where
  expire is not null
  and expire > 0;
```

### Find users belonging to a specific realm

```sql
select
  userid,
  comment
from
  proxmox_user
where
  userid like '%@pam';
```
