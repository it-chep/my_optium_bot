-- +goose Up
-- +goose StatementBegin

-- ТЕРАПИЯ
insert into doctor_messages (scenario_id, next_step, message)
values (4, 12, '🔥Реакция на терапию у пациента {{.FullName}}, {{date_format .BirthDate}}'),
       (4, 18, '🔥Реакция на терапию у пациента {{.FullName}}, {{date_format .BirthDate}}');

insert into admin_messages (scenario_id, next_step, message)
values (4, 3, 'Клиенту ({{.FullName}}, {{date_format .BirthDate}}) нужно помочь сделать заказ');
-- ТЕРАПИЯ

-- МЕТРИКИ
insert into doctor_messages (scenario_id, next_step, message)
values (2, 8, '📈 Метрики заполнены ({{.FullName}}, {{date_format .BirthDate}}). Посмотрите, пожалуйста - {{.MetricsLink}}.

Кратность метрик можно изменить в админке');
-- МЕТРИКИ

-- ВЫВЕДЕНИЕ НА КОНТРОЛЬ
insert into doctor_messages (scenario_id, next_step, message)
values (10, 4, '❗На контроль {{.FullName}}, {{date_format .BirthDate}}'),
       (10, 12, '❗Пациент {{.FullName}}, {{date_format .BirthDate}} отменяет контроль, проверьте причину в чате');

insert into admin_messages (scenario_id, next_step, message)
values (10, 6, '❗На контроль готов {{.FullName}}, {{date_format .BirthDate}}');
-- ВЫВЕДЕНИЕ НА КОНТРОЛЬ


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
