-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email,
  user_image
) VALUES (
  $1, $2, $3, $4 , $5
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
  set username = $2,
      full_name = $3,
      hashed_password = $4,
      user_image = $5
WHERE email = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE username = $1;