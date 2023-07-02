-- +goose Up
-- +goose StatementBegin
create table aliens (
  id serial primary key,
  username text not null,
  email text not null,
  password text not null,
  description text not null,

  created_at timestamp not null default now(),
  updated_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table aliens;
-- +goose StatementEnd
