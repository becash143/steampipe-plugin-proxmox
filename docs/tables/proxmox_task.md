# Table: proxmox_task

Retrieve information about scheduled and historical tasks (e.g. backups, migrations, VM operations) executed across your Proxmox VE cluster.

## Examples

### List recent tasks

```sql
select
  upid,
  node,
  type,
  status,
  start_time
from
  proxmox_task
order by
  start_time desc
limit
  20;
```

### Find failed tasks

```sql
select
  upid,
  node,
  type,
  status,
  start_time
from
  proxmox_task
where
  status != 'OK'
order by
  start_time desc;
```

### Find long-running tasks

```sql
select
  upid,
  node,
  type,
  start_time,
  end_time,
  extract(epoch from (end_time - start_time)) as duration_seconds
from
  proxmox_task
where
  end_time is not null
order by
  duration_seconds desc
limit
  10;
```

### Count tasks by type

```sql
select
  type,
  count(*) as task_count
from
  proxmox_task
group by
  type
order by
  task_count desc;
```

### Find tasks initiated by a specific user

```sql
select
  upid,
  node,
  type,
  status,
  start_time
from
  proxmox_task
where
  "user" = 'root@pam'
order by
  start_time desc;
```

### Find tasks still running

```sql
select
  upid,
  node,
  type,
  start_time
from
  proxmox_task
where
  status = 'running';
```
