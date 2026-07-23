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
  starttime
from
  proxmox_task
order by
  starttime desc
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
  starttime
from
  proxmox_task
where
  status != 'OK'
order by
  starttime desc;
```

### Find long-running tasks

```sql
select
  upid,
  node,
  type,
  starttime,
  endtime,
  extract(epoch from (endtime - starttime)) as duration_seconds
from
  proxmox_task
where
  endtime is not null
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
