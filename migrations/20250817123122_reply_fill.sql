-- +goose Up
-- +goose StatementBegin

-- ТЕРАПИЯ
insert into doctor_messages (scenario_id, next_step, message)
values (4, 12, '🔥Реакция на терапию у пациента {{.FullName}}, {{.BirthDate}}'),
       (4, 18, '🔥Реакция на терапию у пациента {{.FullName}}, {{.BirthDate}}');

insert into admin_messages (scenario_id, next_step, message)
values (4, 3, 'Клиенту ({{.FullName}}, {{.BirthDate}}) нужно помочь сделать заказ');
-- ТЕРАПИЯ

-- МЕТРИКИ
insert into doctor_messages (scenario_id, next_step, message)
values (2, 8, '📈 Метрики заполнены ({{.FullName}}, {{.BirthDate}}). Посмотрите, пожалуйста.

Кратность метрик можно изменить в админке');
-- МЕТРИКИ

-- ВЫВЕДЕНИЕ НА КОНТРОЛЬ
insert into doctor_messages (scenario_id, next_step, message)
values (10, 4, '❗На контроль {{.FullName}}, {{.BirthDate}}'),
       (10, 12, '❗Пациент {{.FullName}}, {{.BirthDate}} отменяет контроль, проверьте причину в чате');

insert into admin_messages (scenario_id, next_step, message)
values (10, 6, '❗На контроль готов {{.FullName}}, {{.BirthDate}}');
-- ВЫВЕДЕНИЕ НА КОНТРОЛЬ


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
