-- +goose Up
-- +goose StatementBegin
ALTER TABLE doctors_scenarios
    ADD COLUMN chat_id bigint;

ALTER TABLE doctors_scenarios
    DROP CONSTRAINT doctors_scenarios_doctor_id_scenario_id_key;

ALTER TABLE doctors_scenarios
    ADD CONSTRAINT doctors_scenarios_doctor_id_scenario_id_chat_id_key
        UNIQUE (doctor_id, scenario_id, chat_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table doctors_scenarios
    drop column chat_id;
-- +goose StatementEnd
