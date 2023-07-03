create table "User"
(
    uid   varchar not null
        primary key
        unique,
    name  varchar,
    email varchar
);

alter table "User"
    owner to table_admin;

create table "Device"
(
    device_id   varchar not null
        constraint device_pk
            primary key,
    name        varchar,
    type        varchar,
    os_name     varchar,
    os_version  varchar,
    app_version varchar,
    uid         varchar
        constraint uid
            references "User"
            on update cascade on delete cascade
);

alter table "Device"
    owner to table_admin;

create table "BatteryLevel"
(
    time           timestamp,
    battery_level  integer not null,
    battery_status varchar not null,
    device_id      varchar
        constraint device
            references "Device"
            on update cascade on delete cascade
);

alter table "BatteryLevel"
    owner to table_admin;

