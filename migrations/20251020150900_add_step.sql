-- +goose Up
-- +goose StatementBegin
insert into scenario_steps (scenario_id, step_order, content, is_final, next_step, next_delay)
values (10, 13, 'Укажите, пожалуйста, дату', true, null, null);

insert into step_buttons (scenario, step, button_text, next_step_order)
values (10, 5, 'Планирую примерно к дате ...', 13);

insert into doctor_messages (scenario_id, next_step, message)
values (10, 99, '❗На контроль {{.FullName}}, {{date_format .BirthDate}}, клиент указал дату планирования {{.Text}}');

insert into admin_messages (scenario_id, next_step, message)
values (10, 99, '❗На контроль ({{.FullName}}, {{date_format .BirthDate}}), клиент указал дату планирования {{.Text}}');

insert into doctor_messages (scenario_id, next_step, message)
values (10, 100, '❗На контроль {{.FullName}}, {{date_format .BirthDate}}, отменяет контроль: {{.Text}}');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
