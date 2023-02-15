create table if not exists "User"
(
    id   varchar not null
        constraint user_pk
            primary key,
    name varchar
);

alter table "User"
    owner to table_admin;

create table if not exists "Device"
(
    id              varchar not null
        constraint device_pk
            primary key,
    name            varchar,
    type            varchar,
    "osName"        varchar,
    "osVersion"     varchar,
    "appVersion"    varchar,
    "batteryLevel"  integer,
    "batteryStatus" varchar,
    "pushToken"     varchar,
    "userId"        varchar
        constraint "userId"
            references "User"
);

alter table "Device"
    owner to table_admin;

create unique index if not exists device_id_uindex
    on "Device" (id);

create unique index if not exists user_id_uindex
    on "User" (id);

