-- +goose Up
-- +goose StatementBegin
insert into posts_themes(name, is_required)
values ('Обязательный контент', true), -- 1
       ('Мотивация', true),            -- 2
       ('Подготовка к новому этапу', true) -- 3
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
