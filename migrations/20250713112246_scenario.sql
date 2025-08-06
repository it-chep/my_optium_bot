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
    step_order  integer not null,      -- порядок
    content     text    not null,
    is_final    boolean default false, -- если тру, то пускаем новый сценарий и в patient_scenarios проставляем complete
    next_delay  interval,
    next_step   int
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

-- связка юзер и сценарий (тут сценарии только друг за другом)
-- если тут нужно чтобы сценарий повторялся и был периодичным, то при завершении этого сценария создаем новую запись в таблице просто

-- чтобы начать новый сценарий будет джоба, которая проверит эту табличку
-- если нет активных сценариев у челика, то назначаем автоматом первый сценарий !completed order by id/created_at
-- джоба должна будет уметь обновить степ и отправить первое сообщение сценария
create table if not exists patient_scenarios
(
    id             serial primary key,
    patient_id     integer   not null,               -- patients(tg_id)
    chat_id        bigint    not null,
    scenario_id    integer   not null,               -- scenarios(id)
    step           integer   not null,               -- scenarios(id)
    answered       boolean   not null default false,
    sent           boolean   not null default false,
    -- когда запланировано. будет отдельная джоба, которая посмотрит это поле, отправит сообщение и сдвинет степ
    -- отдельная джоба нужна на случаи задержки сообщений (например если чел ответил позже, то след сообщение отправим через 3 часа)
    scheduled_time timestamp not null,
    active         boolean   not null default false, -- активен ли сейчас
    repeatable     boolean   not null default false, -- повторяемый ли, если повторяемый, то при завершении будем создавать новую запись в табличку
    completed_at   timestamp                         -- когда завершили
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
