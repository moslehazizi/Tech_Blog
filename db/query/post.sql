-- name: CreatePost :one
INSERT INTO posts (
  post_image,
  title,
  post_category,
  content,
  time_for_read,
  like_number
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
  set post_image = $2,
      title = $3,
      post_category = $4,
      content = $5,
      time_for_read = $6
WHERE id = $1
RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;