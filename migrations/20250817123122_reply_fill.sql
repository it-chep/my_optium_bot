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

-- ПОТЕРЯШКА
insert into doctor_messages (scenario_id, next_step, message)
values (9, 2, '🚷 Клиент {{.FullName}}, {{date_format .BirthDate}} вышел на связь'),
       (9, 3, '🚷 Клиент {{.FullName}}, {{date_format .BirthDate}} вышел на связь'),
       (9, 6, '🚷 Клиент {{.FullName}}, {{date_format .BirthDate}} отложил ведение'),
       (9, 10, ' 🚷Клиент {{.FullName}}, {{date_format .BirthDate}} решил остановить ведение, потому что Достиг цели'),
       (9, 11, ' 🚷Клиент {{.FullName}}, {{date_format .BirthDate}} решил остановить ведение, потому что Нет времени'),
       (9, 12, ' 🚷Клиент {{.FullName}}, {{date_format .BirthDate}} решил остановить ведение, потому что Пока нет финансов на это'),
       (9, 13, ' 🚷Клиент {{.FullName}}, {{date_format .BirthDate}} решил остановить ведение, потому что Были побочные реакции'),
       (9, 14, ' 🚷Клиент {{.FullName}}, {{date_format .BirthDate}} решил остановить ведение, потому что Не получил результата'),
       (9, 15, ' 🚷Клиент {{.FullName}}, {{date_format .BirthDate}} решил остановить ведение. Проверьте причину в чате');
-- ВЫВЕДЕНИЕ НА КОНТРОЛЬ


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
