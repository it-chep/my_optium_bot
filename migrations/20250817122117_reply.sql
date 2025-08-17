-- +goose Up
-- +goose StatementBegin
create table admin_chat
(
    id      serial,
    chat_id bigint not null,
    unique (chat_id)
);

create table admin_messages
(
    id          serial,
    scenario_id bigint  not null,
    next_step   integer not null,
    message     text    not null
);

create table doctor_messages
(
    id          serial,
    scenario_id bigint  not null,
    next_step        integer not null,
    message     text    not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists amin_chat;
drop table if exists amin_messages;
drop table if exists doctor_messages;
-- +goose StatementEnd
