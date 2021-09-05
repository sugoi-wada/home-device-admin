
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table cp_users (
  id serial primary key,
  email varchar not null unique,
  cp_token varchar not null,
  expire_time varchar not null,
  refresh_token varchar not null,
  m_version varchar not null,
  created_at timestamp default null,
  updated_at timestamp default null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table cp_users;
