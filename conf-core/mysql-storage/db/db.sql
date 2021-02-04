create database if not exists conf charset utf8mb4;
use conf;
set global sql_mode = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
create table if not exists profile
(
    id          varchar(64) primary key not null,
    conf_key    varchar(128)  not null,
    conf_value  varchar(2048) not null,
    env         varchar(36)   not null,
    service     varchar(36)   not null,
    description varchar(128),
    status      varchar(10),
    version     int,
    created_by   varchar(36),
    deleted_by   varchar(36),
    modified_by varchar(36),
    created_at   timestamp     null default current_timestamp(),
    deleted_at   timestamp     null default current_timestamp(),
    modified_at timestamp     null default current_timestamp() on update current_timestamp(),

    index conf_key_index (conf_key),
    unique index service_env_key_index (service, env, conf_key)
);

/* 分布式锁 */
create table if not exists profile_lock
(
    name varchar(32) primary key not null,
    owner varchar(32) not null ,
    created_at timestamp null default current_timestamp(),
    expired_at timestamp null default current_timestamp(),

    index expired_index (expired_at)
);