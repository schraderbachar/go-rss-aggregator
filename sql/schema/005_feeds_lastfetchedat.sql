-- +goose Up
alter table feeds add COLUMN last_fetched_at timestamp;

-- +goose Down
alter table feeds drop COLUMN last_fetched_at;