-- +goose Up
create table posts (
    id UUID PRIMARY Key,
    created_at timestamp not null,
    updated_at timestamp not null,
    title text not null,
    description text,
    published_at timestamp not null,
    url text not null unique,
    feed_ID UUID not null references feeds(id) on delete cascade
);

-- +goose Down
drop table users;