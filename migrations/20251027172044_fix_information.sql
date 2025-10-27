-- +goose Up
-- +goose StatementBegin
-- Обновляем контент для существующего шага
UPDATE scenario_steps
SET content = '{{.RequiredPost}}'
WHERE scenario_id = 8
  AND step_order = 3;

-- Временно обновляем step_order чтобы избежать конфликта уникальности
UPDATE scenario_steps
SET step_order = 6,
    next_step  = 7
WHERE scenario_id = 8
  AND step_order = 4;

-- Вставляем новые шаги с правильной последовательностью
INSERT INTO scenario_steps (scenario_id, step_order, content, is_final, next_delay, next_step)
VALUES (8, 4, '{{.AdditionalPost}}', false, '15 seconds', 5),
       (8, 5, '{{.MotivationPost}}', false, '15 seconds', null),
       (8, 7, '{{.SecondPartPost}}', false, '15 seconds', 2);


-- repetitions таблица с количеством повторений сценариев
create table if not exists repetitions
(
    id                bigserial primary key,
    scenario_id       bigint,
    patient_tg_id     bigint,
    repetitions_count bigint,

    unique (patient_tg_id, scenario_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
