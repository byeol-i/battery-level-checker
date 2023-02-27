ALTER SYSTEM SET wal_level = 'hot_standby';
ALTER SYSTEM SET max_wal_senders = 5;
ALTER SYSTEM SET wal_keep_segments = 8;
ALTER SYSTEM SET track_activity_query_size=4096;
ALTER SYSTEM SET log_statement = 'mod';

CREATE USER replicator REPLICATION LOGIN ENCRYPTED PASSWORD 'password';
