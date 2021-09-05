
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table cp_device_states(
  id serial primary key,
  cp_device_id integer references cp_devices on delete cascade unique not null,
  power varchar not null,
  feature varchar not null,
  speed varchar not null,
  temp varchar not null,
  inside_temp varchar not null,
  nanoex varchar not null,
  people varchar not null,
  outside_temp varchar not null,
  pm25 varchar not null,
  on_timer varchar not null,
  off_timer varchar not null,
  vertical_direction varchar not null,
  horizontal_direction varchar not null,
  fast varchar not null,
  econavi varchar not null,
  volume varchar not null,
  display_light varchar not null,
  created_at timestamp default null,
  updated_at timestamp default null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table cp_device_states;