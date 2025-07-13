-- +goose Up
-- +goose StatementBegin

-- основная табличка
create table if not exists scenarios
(
    id          serial primary key,    -- для работы с фронтом
    name        text unique not null,  -- Терапия/Метрики и тд
    description text,                  -- пусть будет
    is_active   boolean  default true, -- чтоб вырубать
    delay       interval
);

-- шаги
create table if not exists scenario_steps
(
    id           serial primary key,
    scenario_id  integer not null,
    step_order   integer not null,     -- порядок
    title        text,                 -- пусть будет скрытый для работы с админкой
    content      bytea   not null,     -- todo: не уверен чет
    content_type int     not null,     -- будем чекать видео/текст/картинка/файл енам
    is_final     boolean default false -- если тру, то пускаем новый сценарий и в patient_scenarios проставляем complete
);

-- детализация шага
create table if not exists step_buttons
(
    id              serial primary key,
    step_id         integer not null, -- scenario_steps(id)
    button_text     text    not null,
    next_step_order integer           -- todo: пока думаю, переход к какому шагу (null - завершение)
    -- doctor_notify   boolean default false -- пздц не факт, ебал рот уведы в лс
);

-- связка юзер и сценарий
create table if not exists patient_scenarios
(
    id              serial primary key,
    patient_id      integer   not null,              -- patients(id)
    scenario_id     integer   not null,              -- scenarios(id)
    scheduled_start timestamp not null,              -- запланировали
    actual_start    timestamp,                       -- по факту старт
    completed_at    timestamp,                       -- когда завершили
--     scenario_on_complete integer not null -- todo: какой сценарий пускать при завершении, мб по другому как то делать хз
    repeatable      boolean   not null default false -- повторяемый ли
);

create table if not exists client_step_answers
(
    id          serial primary key,
    patient_id  integer   not null,-- patients(id)
    scenario_id integer   not null, -- scenarios(id)
    step_id     integer   not null, -- scenario_steps(id)
    answer_text text      not null,
    answer_date timestamp not null default now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists scenarios;
drop table if exists scenario_steps;
drop table if exists step_buttons;
drop table if exists patient_scenarios;
drop table if exists client_step_answers;
-- +goose StatementEnd
