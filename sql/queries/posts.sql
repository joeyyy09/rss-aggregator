-- name: CreatePost :one

INSERT INTO users (id, created_at, updated_at, title, url, feed_id,published_at,description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;