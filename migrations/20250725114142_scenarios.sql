-- +goose Up
-- +goose StatementBegin

-- insert into scenarios (name, description, is_active)
-- values ('init', 'Сценарий запускается, когда врач исполняет команду /init', false);
--
-- insert into scenario_steps (scenario_id, step_order, content, is_final)
-- values
--     (1, 1, 'Укажите Фамилию и Имя клиента', false),
--     (1, 2,'Укажите пол', false),
--     (1, 3,'Укажите дату рождения', false),
--     (1, 4,'Введите ссылку на метрики', true);
--
-- insert into step_buttons (scenario, step, button_text)
-- values
--     (1, 2, 'М'),
--     (1, 2, 'Ж');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
