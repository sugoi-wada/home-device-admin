-- +goose Up
-- +goose StatementBegin
alter table cp_device_states
add column sleep varchar not null default '',
add column dry varchar not null default '',
add column self_clean varchar not null default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table cp_device_states
drop column sleep,
drop column dry,
drop column self_clean;
-- +goose StatementEnd
