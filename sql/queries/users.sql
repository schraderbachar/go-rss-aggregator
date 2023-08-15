-- name: CreateUser :one
insert into users(id,created_at,updated_at,name,api_key)
values ($1,$2,$3,$4,encode(sha256(random()::text::bytea), 'hex')) --this lets the sql do it right here instead of it needing to genrerate its own 
RETURNING *;

-- name: GetUserByAPIKey :one
select * from users where api_key = $1;