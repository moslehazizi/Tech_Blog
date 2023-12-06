-- name: CreateCategory :one
INSERT INTO categories (
  category_name
) VALUES (
  $1
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: UpdateCategory :one
UPDATE categories
  set category_name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;