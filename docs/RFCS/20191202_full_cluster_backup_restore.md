- Feature Name: Full Cluster Backup/Restore
- Status: completed
- Start Date: 2019-12-02
- Authors: Paul Bardea
- RFC PR: #[42887]()
- Cockroach Issue:
  #[44814](https://github.com/cockroachdb/cockroach/issues/44814)


# Summary

Users should be able to `BACKUP` and `RESTORE` all relevant information stored
in their cluster - namely relevant information stored in system tables.

# Motivation

Currently, only user data can be backed up and restored - along with very
limited metadata information (table statistics if requested). There does not
exist a mechanism for a user to easily restore their entire cluster as it
appeared at the time of a backup.

# Guide-level explanation

This RFC builds on the original [Backup &
Restore](https://github.com/cockroachdb/cockroach/blob/master/docs/RFCS/20160720_backup_restore.md)
functionality and extends it to include all logical data stored in the backup.
A new syntax is introduced to perform a full cluster backup and restore:
`BACKUP TO [...]` and `RESTORE FULL CLUSTER FROM [...]`.

Additionally, incremental cluster backups are supported:
```sql
> BACKUP TO 'nodelocal:///cluster-backup/1';
> BACKUP TO 'nodelocal:///cluster-backup/2' INCREMENTAL FROM 'nodelocal:///cluster-backup/1';
```
A user can create an incremental cluster backup, but they must also provide a
full cluster backup and optionally additional incremental backups (as is the
case for non-cluster backups). All listed backups must be full cluster backups.
Incremental cluster backups can be restored in the usual way: `RESTORE FROM
'nodelocal:///cluster-backup/1', 'nodelocal:///cluster-backup/2'`.
Every backup listed must be a full-cluster backup.

A full cluster RESTORE can only be performed in a fresh cluster with no user
data. Some of the data in the system tables may be set (for example the
`cluster.organization` setting must be set in order to even use this feature).
However, it should be noted that this data will be modified by performing a
full cluster RESTORE.

A full cluster BACKUP/RESTORE could be thought of as performing the following steps:
```sql
/* Full Cluster Backup
There are no semantics to restore all user tables. Assume all user databases are: database_a, database_b, [...].
/* Current backup also does not support backing up entire databases and individual tables, but only a subset of system tables should be backed up. */
BACKUP DATABASE database_a, database_b, [...], system TO 'nodelocal:///cluster-backup/1';

/* Full Cluster Restore */
CREATE DATABASE crdb_system_temporary;
RESTORE system.* FROM 'nodelocal:///cluster-backup/1' WITH into_db='crdb_system_temporary';

/* Restore the user data. */
RESTORE DATABASE database_a, database_b, [...] FROM 'nodelocal:///full-cluster-backup';

/* Restore the system tables. */
BEGIN;
DELETE FROM system.users WHERE true;
INSERT INTO system.users (SELECT * FROM crdb_system_temporary.zones);
COMMIT;

BEGIN;
DELETE FROM system.settings WHERE true;
INSERT INTO system.settings (SELECT * FROM crdb_system_temporary.zones);
COMMIT;

[...]
```

Not all system tables should be included in a backup since some information
relates to the physical properties of a cluster. The existing system tables
have been audited below. New system tables will need to add themselves to the
list of system tables that should be included in a backup. This will initially
be a list of names of system tables maintained inside the `backupccl` package.

### System Tables
| Table Name | Description | Included | Notes |
|---|---|---|---|
| namespace | Provides relationship between parentID <-> descriptor name <-> descriptor ID| No | This information should be generated by the restoring cluster |
| descriptor | Maps ID <-> Descriptor Proto| No | New descriptors should be made for every RESTOREd table. |
| users | Stores the users in the table. | Yes | |
| zones | Stores the zone config protos | Yes | |
| settings | Stores all the cluster settings | Yes | |
| leases | Table leases | No | Leases held in the old cluster are no longer relevant. |
| eventlog | A log of a variety of events (schema changes, node additions, etc..) | No | Most events are not node-specific and would be useful to backup. This may produce confusing output if restored into a cluster with a different number of nodes. See Future work. |
| rangelog | Range level events. | No | Ranges on the old and new cluster will not match. |
| ui | A set of KV pairs used by the UI | Yes | |
| jobs | A list of all jobs that are running or have run. | Yes | |
| web\_sessions |  | No | This could eventually be moved into the backup. Unclear. |
| table\_statistics | | Yes | This information is currently backed up in the BACKUP manifest to BACKUP and RESTORE table statistics on a per-table level. |
| locations | Stores information about the localities. | Yes |  |
| role\_members | Contains role-role and user-role relationships for RBAC | Yes |  |
| comments | Stores up to 1 string comment per object ID | Yes | |
| replication\_* | | No | Replication stats should be regenerated when the data is RESTORED. |
| reports_meta | | No | " |
| protectedts\_* | As proposed by the [protected timestamp RFC](https://github.com/cockroachdb/cockroach/blob/master/docs/RFCS/20191009_gc_protected_timestamps.md) | No | Restore only restores a snapshot of the data in the backup, not the entire MVCC history. |

There is no information in the system ranges that should be included in a
CLUSTER backup since they all relate to properties of the ranges/nodes.

# Reference-level explanation

This RFC assumes that a cluster restore will occur on a fresh cluster with no
user data. This allows the data to be restored _exactly_ as it appeared in the
backup. Otherwise, it will be necessary to re-key the user tables as well as
the **content** inside the system tables which reference table IDs/KV spans
(such as zone configs). Note: this implies that behavior resulting in
interactions with the restoring cluster is undefined until the restore
succeeds. This may be extended in the future, see the Future work section for
more details.

Additionally, incremental cluster backups and restoration is supported using
the same syntax as the existing `BACKUP`. In addition to checking that the
previous backups cover the necessary span of keybase and time, a check must be
added `backupPlanHook` to verify that every backup that this incremental backup
builds upon are also cluster backups. Additionally, full cluster restore should
only be permitted on Therefore it is necessary to add a flag in the backup
manifest (`BackupDescriptor`) indicating whether or not a given backup is a
full cluster backup or not. The primary reason for this bit is to ensure that
full cluster restore can only restore full cluster backup files.

## Interfaces

Users will mainly interact with this new feature through the new syntax
introduced: `BACKUP FULL CLUSTER TO [...]` and `RESTORE FULL CLUSTER FROM
[...]`.

Additionally, incremental cluster backups are supported:
```sql
> BACKUP FULL CLUSTER TO 'nodelocal:///cluster-backup/1';
> BACKUP FULL CLUSTER TO 'nodelocal:///cluster-backup/2' INCREMENTAL FROM 'nodelocal:///cluster-backup/1';
```

These backups can be restored: `RESTORE FULL CLUSTER FROM
'nodelocal:///cluster-backup/1', 'nodelocal:///cluster-backup/2'`. Every backup
listed must be a full-cluster backup.

This new syntax introduces a new target "FULL CLUSTER" which can be used
instead of specifying particular databases/tables to be restored. Replacing the
targets for the new target (FULL CLUSTER), should not result in any UX
surprises.

A user can then examine a full cluster backup using the `SHOW BACKUP` command
(`start_time` and `end_time` are omitted from this example for brevity):
```sql
root@:26258/default_db> SHOW BACKUP 'nodelocal://1/full-cluster';
  database_name | table_name | size_bytes | rows | is_full_cluster |
+---------------+------------+------------+------+-----------------+
  some_user_db  | foo        |          0 |    0 |      true       |
  system        | zones      |        252 |    0 |      true       |
  system        | users      |         99 |    0 |      true       |
  ...
```

This command shows the user what type of metadata is stored in the backup.
Since users must specify only full cluster backups to build incremental
backups, this allows users to inspect a backup to check what cluster
information is stored.

With regards to user-visible errors introduced by this feature, users can
expect to see an error when:
- They create an full cluster incremental backup on top of a non-full cluster
  backup.
- They perform a full cluster restore in a cluster with existing user data
  (there may be table/database ID collisions, which will not be handled). As
  described, a check will be performed ensuring that no user tables/databases
  have been created (practically, this means ensuring that no descriptors
  should exist with an ID greater than or equal to
  `MinNonPredefinedUserDescID`).
- They attempt to perform a full cluster restore from a non-full cluster
  backup.

Note: it is expected that users will be able to perform a non-full cluster
RESTORE on a full-cluster BACKUP.

## Detailed design

## Backup

The first difference between a full cluster backup and a regular (non-full
cluster) backup is that a full cluster backup includes all user tables in the
backup. This can be accomplished by including all tables -- as defined by
enumerating the descriptors table -- except for the set explicitly excluded as
defined above.

Additionally, all OFFLINE tables need to be included in a BACKUP (they are not
today). This is used to ensure that the in-progress jobs may be able to
continue after a full cluster restoration. See the Jobs section below.

Finally, the backup manifest (`BackupDescriptor` protobuf) needs to be augmented
with an enum specifying the amount of cluster information stored in the backup.
An enum `DescriptorCoverage` will be added to the `BackupDescriptor` and will
have options: `RequestedDescriptors`, which is the default and is what existing
backups will have going forward, and `AllDescriptors` for full cluster backup.
This enum is required to prevent a full cluster restore being performed from a
non-full cluster backup file. In particular, this requirement exists because
full cluster RESTORE guarantees that the entire cluster has been RESTOREd (so we
need the entire cluster to be in the backup file).

## Restore

Upon a full cluster restore, the order in which data is restored becomes
relevant. In particular, `system.zones` must be restored prior to restoring the
user data in order to ensure that the user data is placed in the appropriate
locality if appropriate. The user data will then be restored, and finally the
rest of the system tables.

First, a check is performed to ensure that no user data exists in this cluster.
This is achieved by ensuring that no descriptors exist with ID greater than or
equal to `MinNonPredefinedUserDescID`. Then the `DescIDGenerator` needs to be
restored. This key is used to determine that value of the next descriptor ID
(such as during the creation of a table or database). This check would also
ensure that no other full cluster restores are in progress, as the full cluster
restore would create a `crdb_system_temporary` table in the user database
space. It is incremented whenever a descriptor is created. Let `MaxID` by the
maximum descriptor ID found in the backup, then the `DescIDGenerator` should be
set to `MaxID + 1` so that new descriptors can be created after the restore
with correct IDs. 

System tables cannot be restored in the same way as user data tables are since
they occupy a fixed keyspace (and thus cannot be re-keyed as we do today for
new tables). First we restore the system tables into a temporary database. The
`DescIDGenerator` key must be updated prior to creating this temporary table to
ensure no conflicts with a user table that needs to be restored (and thus the
ID of this table will be `MaxID + 1` and the `DescIDGenerator` will be
incremented again).
```sql
CREATE DATABASE crdb_system_temporary;
RESTORE system.* FROM 'nodelocal://1/full-backup/1' WITH into_db='crdb_system_temporary';
```

In an internal executor execute:
```sql
BEGIN;
DELETE FROM system.zones WHERE true;
INSERT INTO system.zones (SELECT * FROM crdb_system_temporary.zones);
COMMIT;
```

Before restoring the user data, we need to ensure that all the user tables and
database descriptors are created with the same ID as they have in the backup.
This differs from the current implementation which generates a new ID for these
items. This allows for a potential future optimization to skip the no-op
rekeying. User tables can then be restored normally.

Finally, to restore the remaining of the system tables, perform a transaction
similar to the one listed above, but rather than only restoring the zones table,
restore the rest of the system tables. It is preferable to restore all of the
remaining system tables in one transaction in order to ensure atomicity across
the restoration of all the system tables.  However, there may be a limitation
based on the maximum transaction size, in which case the possibility of
restoring the system tables 1 by 1 could be investigated. However, the maximum
size of a transaction is quite large and is _not_ expected to cause issue.

Finally, the temporary `crdb_system_temporary` database is deleted.

### Jobs

During a cluster backup, a job may be in progress. The state for these jobs
should persisted in the user-data and in the `system.jobs` table. These jobs
will be restored into a running state and nobody will have a lease on this job.
This job should be adopted and the continued.  For the job to be able to be
continued, all OFFLINE tables need to be included in the BACKUP.

## Locality Awareness

The current implementation of locality-aware BACKUPs should continue to work
with cluster backup without further work. BACKUP for the system tables will
operate just as the user-data tables and the relevant lease-holders will backup
to the appropriate locality.

## Failure Modes

### General Restore Failure/Cancellation

The happy path for a full cluster backup is when the restore is started and all
nodes remain available until the restore is complete.

Non-cluster restore creates the tables for the user-data tables at the start of
the restore. These tables are in an OFFLINE state - inaccessible to the
user[1]. If there is a failure during the restore, and these tables are marked
as DROP and will be removed. Full cluster restore can recover the user-data
tables that it restored this way as well. The difficulty lies in handling
system tables that it has already restored. This will likely be only the
`system.zones` table since the remainder of the system tables will be restored
in a single transaction near the end of the job, however the general case is
considered.

Since full cluster backup must have been run on a fresh cluster, the first
iteration of full cluster restore could require the cluster to be destroyed if
the restore fails. This can likely be improved as detailed in the Future Work
section.


# Alternatives Considered

## System Table Restorations

### AddSSTable and TimeBoundedDelete

One reasonable question is why doesn't CockroachDB load the system data the
same way as the user-data tables. One difficulty that this would present is
that user-data restoration happens on new tables, but the system tables in the
new cluster already have data in them. This method directly ingests the
SSTables for the system tables spans, then issues a `TimeBoundedDelete`. This
does not yet exist, but can be implemented by leveraging
`engine.MVCCClearTimeRange`, similarly to the `batcheval.RevertRange` command.
This leaves the possibility of having a potentially dirty state in the system
tables. Additionally, the keys in the SSTs would need to have their timestamp
updated to some time greater than the time of the start of the restore.

The reason we take an approach of loading SSTs directly into the storage layer
for user data is that we typically expect a large volume of data. Additionally,
we can ensure that this data is not needed or accessed by the user while it
is being loaded. Since the size of the system tables is expected to be much
smaller than the size of the user-data tables, there are no advantages to
this approach and it is more complex.

## Cluster Info Metadata

### Only Look at Backup Contents
Additionally, instead of marking a particular backup as "full cluster", the
system tables that it holds could be examined. This would allow for previous
backups that included the system table information to be restored via a full
cluster backup. One problem with this approach is that if new system tables are
added to the list of expected system tables, a mapping between version numbers
and which tables are expected to be included in that version needs to be
maintained. This problem is avoided by marking backups as full cluster since we
can assume that all system tables included in those backups are safe to restore
(and override the existing ones).

# Future Work

- As mentioned, it may be possible to enable cluster restoration on a non-new
  cluster - however this does raise further complications. Since it seems that
  the vast majority of uses cases for this feature are to restore a cluster
  exactly how it was in the backup, there is little motivation to generalize
  cluster restoration in this way. In particular, this would require the
  contents of the system tables to be re-keyes (in addition to the user-data
  KVs themselves). This would require each system table to provide a way to map
  each of their rows to an updated row based on a re-keyer.

- One large remaining piece of work is how to handle the case where the
  metadata of the restoring cluster does not match that of the one in the
  backup. For example, if the cluster on which the BACKUP was performed as a
  given set of localities which do not exist on the cluster that is being
  restored, there is currently no way for the user to map the localities from
  the backed up cluster to the values they should be changed to in the
  restoring cluster. Currently since all BACKUP and RESTORE interactions
  happens at the CLI, a major difficulty is providing a powerful enough
  interface for the user to provide these mappings.

- Include `system.eventlog` in a full cluster backup. One reason for not doing
  this initially is that some event logs may be non-sensical if the table is
  restored in a cluster with a different number of nodes.

- Additionally, one further improvement would be allow the restoration of a set
  of tables/databases with their respective configuration. This requires that
  the RESTORE process find which rows in the system tables are applicable to
  the given database/table. This also implies that we'd need to add the ability
  to rewrite values in other tables. This is out of scope of this RFC.

- A potential improvement to ensure that the cluster is in a fresh state would
  be to mark a cluster for restoration at creation time (similar to `cockroach
  init`). This would also prevent any operations to interfere with the restore.

- A more graceful failure mode could be implemented which ensure's that the
  cluster's state is guaranteed to be healthy in the case of a failed full
  cluster RESTORE.

- The initial implementation will not consider what happens if there is a
  failure in the middle of the backup. It will clean up the data following the
  normal backup procedures. In the case that there is a failure while updating
  the system tables, the cluster should be started up again. Since we enforce
  that the cluster we are restoring to has no user data, this is acceptable.

# Drawbacks

Due to the restriction that full cluster backup can only be performed on a
newly created cluster, there may be some user surprised when trying to perform
a full cluster restore when this assumption is violated.



[1] Offline tables can however be references when setting zone configs. See
#40285.
