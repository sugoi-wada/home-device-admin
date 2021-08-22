
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table cp_devices (
  id serial not null,
  gateway_id varchar not null,
  auth varchar not null,
  device_id varchar not null,
  nickname varchar not null,
  created_at timestamp default null,
  updated_at timestamp default null,
  primary key(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table cp_devices;
