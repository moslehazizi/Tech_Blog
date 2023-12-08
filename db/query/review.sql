-- name: CreateReview :one
INSERT INTO reviews (
  reviewer,
  review_content,
  post,
  star_degree
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetReview :one
SELECT * FROM reviews
WHERE id = $1 LIMIT 1;


-- name: ListReviews :many
SELECT * FROM reviews
WHERE 
    post = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: DeleteReview :exec
DELETE FROM reviews
WHERE id = $1;