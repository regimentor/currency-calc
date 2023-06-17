drop table if exists api_logs;
drop table if exists users;
drop table if exists currencies;

create table users
(
    id      serial primary key,
    api_key text not null
);


create table currencies
(
    id    serial primary key,
    slug  text  not null,
    value decimal not null,
    date  date  not null,
    base  text  not null
);


create table api_logs
(
    id           serial primary key,
    user_id      integer not null,
    request_type int     not null,
    request_time date    not null,

    foreign key (user_id) references users (id)
);