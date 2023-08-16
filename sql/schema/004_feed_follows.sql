-- +goose Up
create table feed_follows (
    id UUID primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    user_id UUID NOT NULL REFERENCES users(id) on delete cascade,
    feed_id UUID NOT NULL REFERENCES feeds(id) on delete cascade,
    UNIQUE(user_id,feed_id)
);

-- +goose Down
drop table feed_follows;