-- +goose Up
create table if not exists users (
    id serial primary key,
    name varchar not null,
    email varchar not null,
    password varchar not null,
    role varchar not null,
    created_at timestamp not null,
    updated_at timestamp not null
);

-- +goose Down
drop table users;
