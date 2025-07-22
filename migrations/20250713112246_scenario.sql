-- +goose Up
-- +goose StatementBegin

-- основная табличка
create table if not exists scenarios
(
    id          serial primary key,   -- для работы с фронтом
    name        text unique not null, -- Терапия/Метрики и тд
    description text,                 -- пусть будет
    is_active   boolean default true, -- чтоб вырубать
    delay       interval
);

-- шаги
create table if not exists scenario_steps
(
    id          serial primary key,
    scenario_id integer not null,
    step_order  integer not null,     -- порядок
    content     text    not null,
    is_final    boolean default false -- если тру, то пускаем новый сценарий и в patient_scenarios проставляем complete
--     title        text,                 -- пусть будет скрытый для работы с админкой
--     content_type int     not null,     -- будем чекать видео/текст/картинка/файл енам
);

-- детализация шага
create table if not exists step_buttons
(
    id              serial primary key,
    scenario        integer not null, -- scenarios(id)
    step            integer not null, -- scenario_steps(step_order)
    button_text     text    not null,
    next_step_order integer           -- todo: пока думаю, переход к какому шагу (null - завершение)
    -- doctor_notify   boolean default false -- пздц не факт, ебал рот уведы в лс
);

-- связка юзер и сценарий
create table if not exists patient_scenarios
(
    id              serial primary key,
    patient_id      integer   not null,               -- patients(tg_id)
    scenario_id     integer   not null,               -- scenarios(id)
    step            integer   not null,               -- scenarios(id)
    scheduled_start timestamp not null,               -- запланировали
    actual_start    timestamp,                        -- по факту старт
    completed_at    timestamp,                        -- когда завершили
--     scenario_on_complete integer not null -- todo: какой сценарий пускать при завершении, мб по другому как то делать хз
    repeatable      boolean   not null default false, -- повторяемый ли
    unique (patient_id, scenario_id)
);

create table if not exists patient_step_answers
(
    id          serial primary key,
    patient_id  integer   not null,-- patients(id)
    scenario_id integer   not null, -- scenarios(id)
    step_id     integer   not null, -- scenario_steps(id)
    answer_text text      not null,
    answer_date timestamp not null default now()
);

-- связка юзер и сценарий
create table if not exists doctors_scenarios
(
    id           serial primary key,
    doctor_id    integer not null, -- doctors(tg_id)
    scenario_id  integer not null, -- scenarios(id)
    step         integer not null, -- scenarios(id)
    completed_at timestamp,        -- когда завершили
    unique (doctor_id, scenario_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists scenarios;
drop table if exists scenario_steps;
drop table if exists step_buttons;
drop table if exists patient_scenarios;
drop table if exists client_step_answers;
drop table if exists doctors_scenarios;
-- +goose StatementEnd
