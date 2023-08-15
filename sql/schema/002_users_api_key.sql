-- +goose Up
alter table users add COLUMN api_key VARCHAR UNIQUE NOT NULL DEFAULT ( -- if didn't set default we would have a problem with the ones that are already there
    encode(sha256(random()::text::bytea), 'hex')
);

-- +goose Down
alter table users drop COLUMN api_key