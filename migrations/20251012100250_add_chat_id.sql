-- +goose Up
-- +goose StatementBegin
alter table doctors_scenarios
    add column chat_id bigint;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table doctors_scenarios
    drop column chat_id;
-- +goose StatementEnd
