// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: post.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  post_image,
  title,
  post_category,
  content,
  time_for_read,
  like_number
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, post_image, title, post_category, content, time_for_read, like_number, created_at
`

type CreatePostParams struct {
	PostImage    []byte `json:"post_image"`
	Title        string `json:"title"`
	PostCategory int64  `json:"post_category"`
	Content      string `json:"content"`
	TimeForRead  int32  `json:"time_for_read"`
	LikeNumber   int64  `json:"like_number"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.PostImage,
		arg.Title,
		arg.PostCategory,
		arg.Content,
		arg.TimeForRead,
		arg.LikeNumber,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.PostImage,
		&i.Title,
		&i.PostCategory,
		&i.Content,
		&i.TimeForRead,
		&i.LikeNumber,
		&i.CreatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id, post_image, title, post_category, content, time_for_read, like_number, created_at FROM posts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.PostImage,
		&i.Title,
		&i.PostCategory,
		&i.Content,
		&i.TimeForRead,
		&i.LikeNumber,
		&i.CreatedAt,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, post_image, title, post_category, content, time_for_read, like_number, created_at FROM posts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.PostImage,
			&i.Title,
			&i.PostCategory,
			&i.Content,
			&i.TimeForRead,
			&i.LikeNumber,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostsByCategory = `-- name: ListPostsByCategory :many
SELECT id, post_image, title, post_category, content, time_for_read, like_number, created_at FROM posts
WHERE 
  post_category = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListPostsByCategoryParams struct {
	PostCategory int64 `json:"post_category"`
	Limit        int32 `json:"limit"`
	Offset       int32 `json:"offset"`
}

func (q *Queries) ListPostsByCategory(ctx context.Context, arg ListPostsByCategoryParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPostsByCategory, arg.PostCategory, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.PostImage,
			&i.Title,
			&i.PostCategory,
			&i.Content,
			&i.TimeForRead,
			&i.LikeNumber,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
  set post_image = $2,
      title = $3,
      post_category = $4,
      content = $5,
      time_for_read = $6
WHERE id = $1
RETURNING id, post_image, title, post_category, content, time_for_read, like_number, created_at
`

type UpdatePostParams struct {
	ID           int64  `json:"id"`
	PostImage    []byte `json:"post_image"`
	Title        string `json:"title"`
	PostCategory int64  `json:"post_category"`
	Content      string `json:"content"`
	TimeForRead  int32  `json:"time_for_read"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost,
		arg.ID,
		arg.PostImage,
		arg.Title,
		arg.PostCategory,
		arg.Content,
		arg.TimeForRead,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.PostImage,
		&i.Title,
		&i.PostCategory,
		&i.Content,
		&i.TimeForRead,
		&i.LikeNumber,
		&i.CreatedAt,
	)
	return i, err
}
