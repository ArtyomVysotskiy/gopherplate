create schema if not exists test;

create table if not exists test.test
(
    id serial primary key,
    name text
);

