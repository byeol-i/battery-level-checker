create table if not exists "User"
(
    user_id varchar not null
        primary key,
    name    varchar,
    email   varchar
);

alter table "User"
    owner to table_admin;

create table if not exists "Device"
(
    device_id   serial
        constraint device_pk
            primary key,
    name        varchar,
    type        varchar,
    os_name     varchar,
    os_version  varchar,
    app_version varchar,
    user_id     varchar
        constraint user_id
            references "User"
            on update cascade on delete cascade
);

alter table "Device"
    owner to table_admin;

create table if not exists "BatteryLevel"
(
    time           timestamp,
    battery_level  integer not null,
    battery_status varchar not null,
    device_id      integer
        constraint device
            references "Device"
            on update cascade on delete cascade
);

alter table "BatteryLevel"
    owner to table_admin;

