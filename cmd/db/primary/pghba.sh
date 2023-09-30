#!/bin/sh

PG_HBA_CONF="${PGDATA}/pg_hba.conf"

grep replicator "$PG_HBA_CONF" || (echo "host replication     replicator      0.0.0.0/0            md5" >> "$PG_HBA_CONF")
grep replicator "$PG_HBA_CONF" || (echo "host replication     replicator      ::/0                md5" >> "$PG_HBA_CONF")

echo "pg_hba.conf is ready"
