-- +goose Up
-- +goose StatementBegin
create table if not exists patients
(
    id           bigserial not null,
    tg_id        bigint primary key,
    created_at   timestamp not null default now(),
    full_name    text,
    sex          int,
    birth_date   date,
    metrics_link text
);

create table if not exists doctors
(
    id          serial             not null,
    tg_id       bigint primary key not null,
    full_name   text               not null,
    tg_username text               not null,
    is_active   boolean default true
);

-- уточнить точно ли может связь или 1к1
create table if not exists patient_doctor
(
    doctor_tg  bigint,
    patient_tg bigint,
    chat_id    bigint,
    unique (doctor_tg, patient_tg)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists patients;
drop table if exists doctors;
drop table if exists patient_doctor;
-- +goose StatementEnd
