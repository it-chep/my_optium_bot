-- +goose Up
-- +goose StatementBegin

-- Рассылки
create table if not exists newsletters
(
    id               bigserial primary key,
    recipients_count int,                              -- количество пользователей
    text             text      not null,               -- текст рассылки
    users_lists      int[],                            -- списки пользователей
    users_ids        bigint[],                         -- id пользователей, которым ушла рассылка
    media_id         varchar(255),                     -- медиа ID
    content_type_id  integer,
    created_at       timestamp not null default now(), -- дата создания
    sent_at          timestamp,                        -- дата отправки
    name             varchar(255),                     -- Название рассылки
    status_id        int                               -- статус рассылки
    -- отложенная отправка ?????
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists newsletters;
drop table if exists newsletters_status;
-- +goose StatementEnd
