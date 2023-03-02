#!/bin/sh

# TODO: Lame attempt to make this idempotent

if [ ! -f ${PGDATA}/recovery.conf ]; then
  #Danger Will Robinson, Danger!
  gosu postgres pg_ctl -D "$PGDATA" -m fast -w stop
  sleep 1
  rm -rf ${PGDATA}/*

  PGPASSWORD="password" pg_basebackup -h master -p 5432 -D ${PGDATA} -U replicator -X stream -v

  # TODO: Do this with ALTER SYSTEM
  cat >> "${PGDATA}/postgresql.conf" <<EOF
  wal_level = hot_standby
  max_replication_slots = 10
  max_wal_senders = 3
  wal_keep_segments = 8
  hot_standby = on
EOF

  cat >> "${PGDATA}/recovery.conf" <<EOF
  standby_mode = 'on'
  primary_conninfo = 'host=master port=5432 user=replicator password=password'
  trigger_file = '/tmp/postgresql.trigger'
  #primary_slot_name = 'replica1'
EOF

  chown -R postgres ${PGDATA}

  gosu postgres pg_ctl -D "$PGDATA" \
    -o "-c listen_addresses='localhost'" \
    -w start
fi
