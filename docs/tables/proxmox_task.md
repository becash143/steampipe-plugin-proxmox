---
title: "Steampipe Table: proxmox_task - Query Proxmox VE Tasks using SQL"
description: "Allows users to query scheduled and historical tasks executed across a Proxmox VE cluster, such as backups, migrations, and VM operations."
folder: "Task"
---

# Table: proxmox_task - Query Proxmox VE Tasks using SQL

Retrieve information about scheduled and historical tasks (e.g. backups, migrations, VM operations) executed across your Proxmox VE cluster.

## Table Usage Guide

The `proxmox_task` table in Steampipe provides you with information about tasks executed across a Proxmox VE cluster. This table allows you, as a systems administrator, to query task-specific details, including task type, status, start and end time, and the user who initiated the task. You can utilize this table to gather insights on cluster activity, such as identifying failed tasks, finding long-running operations, or auditing which user initiated a given task.

**Note:** Proxmox's `/nodes/{node}/tasks` endpoint returns only the ~50 most recent tasks per node by default. This table currently reflects that same limit — older task history beyond the most recent ~50 per node will not appear in query results, even without a `limit` clause in your SQL. Paging through `start`/`limit` to retrieve full history is planned but not yet implemented, pending changes to the underlying API client.

## Examples

### List recent tasks

```sql+postgres
select
  up_id,
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

```sql+sqlite
select
  up_id,
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

```sql+postgres
select
  up_id,
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

```sql+sqlite
select
  up_id,
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

```sql+postgres
select
  up_id,
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

```sql+sqlite
select
  up_id,
  node,
  type,
  start_time,
  end_time,
  (julianday(end_time) - julianday(start_time)) * 86400 as duration_seconds
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

```sql+postgres
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

```sql+sqlite
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

```sql+postgres
select
  up_id,
  node,
  type,
  status,
  start_time
from
  proxmox_task
where
  user_id = 'root@pam'
order by
  start_time desc;
```

```sql+sqlite
select
  up_id,
  node,
  type,
  status,
  start_time
from
  proxmox_task
where
  user_id = 'root@pam'
order by
  start_time desc;
```

### Find tasks still running

```sql+postgres
select
  up_id,
  node,
  type,
  start_time
from
  proxmox_task
where
  status = 'running';
```

```sql+sqlite
select
  up_id,
  node,
  type,
  start_time
from
  proxmox_task
where
  status = 'running';
```
