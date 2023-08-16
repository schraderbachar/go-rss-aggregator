-- name: CreateFeed :one
insert into feeds (id,created_at,updated_at,name,url,user_id)
values ($1,$2,$3,$4,$5,$6)
RETURNING *;


-- user can get all feeds in db

-- name: GetFeeds :many
select * from feeds;

-- name: GetNextFeedsToFetch :many
select * from feeds order by last_fetched_at asc nulls first limit $1;

-- name: MarkFeedAsFetched :one
update feeds set last_fetched_at = NOW(), updated_at = NOW() where id = $1 
RETURNING * ;