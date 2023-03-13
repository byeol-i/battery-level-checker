create table "User"
(
    id   varchar not null
        constraint user_pk
            primary key,
    name varchar
);

alter table "User"
    owner to table_admin;

create unique index user_id_uindex
    on "User" (id);

create table "Device"
(
    id           varchar not null
        constraint device_pk
            primary key,
    name         varchar,
    type         varchar,
    "osName"     varchar,
    "osVersion"  varchar,
    "appVersion" varchar,
    "userId"     varchar
        constraint "userId"
            references "User"
            on update cascade on delete cascade
);

alter table "Device"
    owner to table_admin;

create unique index device_id_uindex
    on "Device" (id);

create table "BatteryLevel"
(
    time            timestamp,
    "batteryLevel"  integer not null,
    "batteryStatus" varchar not null,
    "deviceId"      varchar
        constraint device
            references "Device"
            on update cascade on delete cascade
);

alter table "BatteryLevel"
    owner to table_admin;

