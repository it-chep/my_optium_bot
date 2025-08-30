-- +goose Up
-- +goose StatementBegin

-- Списки пользователей
create table if not exists user_lists
(
    id   bigserial primary key,
    name varchar(255) not null unique -- название списка
);

-- M2M таблица пользователь - список пользователей
create table if not exists users_lists
(
    id      bigserial primary key,
    user_id bigint,
    list_id bigint,

    unique (user_id, list_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists users_lists;
drop table if exists user_lists;
-- +goose StatementEnd
