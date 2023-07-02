-- +goose Up
-- +goose StatementBegin
create table relationships (
    id serial primary key,
    alien_id integer not null,
    target_id integer not null,
    status integer not null,

    created_at timestamp not null default now(),
    updated_at timestamp not null default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table relationships;
-- +goose StatementEnd
