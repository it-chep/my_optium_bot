-- +goose Up
-- +goose StatementBegin

-- Темы постов
create table if not exists posts_themes
(
    id          bigserial,          -- id темы в системе
    name        text not null,      -- Название темы (мотивация, обязательная тема)
    is_required bool default false, -- флаг обязательная ли тема
    theme_order int                 -- порядковый номер темы
);

-- Посты сценария информации
create table if not exists information_posts
(
    id              bigserial,       -- id поста в системе
    name            text   not null, -- название поста для отображения в админке
    posts_theme_id  bigint not null, -- id темы, к которой относится пост
    order_in_theme  int    not null, -- порядковый номер поста в теме
    media_id        text,            -- медиа файл
    content_type_id int,             -- тип медиафайла
    post_text       text   not null  -- текст поста
);

-- Посты, которые получил пользователь
create table if not exists patient_posts
(
    id          bigserial,          -- id в системе
    patient_id  bigint not null,    -- id пациента
    post_id     bigint not null,    -- id поста
    is_received bool default false, -- флаг получил ли пользователь данный пост
    sent_time   timestamp,          -- время отправки

    unique (patient_id, post_id)    -- констрейнт который гарантирует нам, что пользователь должен получить пост только 1 раз
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists posts_themes;
drop table if exists information_posts;
drop table if exists patient_posts;
-- +goose StatementEnd
