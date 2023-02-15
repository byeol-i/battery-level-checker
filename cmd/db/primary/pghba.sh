#!/bin/sh

grep replicator "${PGDATA}/pg_hba.conf" || (echo "host replication     replicator      0.0.0.0/0            md5" >> "${PGDATA}/pg_hba.conf")
grep replicator "${PGDATA}/pg_hba.conf" || (echo "host replication     replicator      ::/0                md5" >> "${PGDATA}/pg_hba.conf")