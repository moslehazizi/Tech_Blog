package db

import (
	"Tech_Blog/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomPost(t *testing.T, category Category) Post {
	arg := CreatePostParams{
		PostImage:    util.RandomBytes(10),
		Title:        util.RandomName(),
		PostCategory: category.ID,
		Content:      util.RandomName(),
		TimeForRead:  util.RandomInt32(1, 5),
		LikeNumber:   util.RandomInt(5, 20),
	}
	post, err := testQueries.CreatePost(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post)
	require.Equal(t, arg.PostImage, post.PostImage)
	require.Equal(t, arg.Title, post.Title)
	require.Equal(t, arg.PostCategory, post.PostCategory)
	require.Equal(t, arg.Content, post.Content)
	require.Equal(t, arg.TimeForRead, post.TimeForRead)
	require.Equal(t, arg.LikeNumber, post.LikeNumber)
	require.NotZero(t, post.ID)

	return post
}

func TestCreatePost(t *testing.T) {
	category := createRandomCategory(t)
	createRandomPost(t, category)
}

func TestGetPost(t *testing.T) {
	category := createRandomCategory(t)
	post_1 := createRandomPost(t, category)
	post_2, err := testQueries.GetPost(context.Background(), post_1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, post_2)

	require.Equal(t, post_1.ID, post_2.ID)
	require.Equal(t, post_1.PostImage, post_2.PostImage)
	require.Equal(t, post_1.Title, post_2.Title)
	require.Equal(t, post_1.PostCategory, post_2.PostCategory)
	require.Equal(t, post_1.Content, post_2.Content)
	require.Equal(t, post_1.TimeForRead, post_2.TimeForRead)
	require.Equal(t, post_1.LikeNumber, post_2.LikeNumber)
}

func TestListPost(t *testing.T) {
	category := createRandomCategory(t)
	for i := 0; i < 10; i++ {
		createRandomPost(t, category)
	}
	arg := ListPostsParams{
		Limit:  5,
		Offset: 5,
	}

	posts, err := testQueries.ListPosts(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, posts, 5)

	for _, post := range posts {		
		require.NotEmpty(t, post)
	}
}

func TestUpdate(t *testing.T) {
	category := createRandomCategory(t)
	post_1 := createRandomPost(t, category)
	category_2 := createRandomCategory(t)
	arg := UpdatePostParams{
		ID:           post_1.ID,
		PostImage:    util.RandomBytes(20),
		Title:        util.RandomName(),
		PostCategory: category_2.ID,
		Content:      util.RandomName(),
		TimeForRead:  util.RandomInt32(6, 8),
	}
	post_2, err := testQueries.UpdatePost(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post_2)
	require.Equal(t, post_1.ID, post_2.ID)
	require.NotEqual(t, post_1.PostImage, post_2.PostImage)
	require.NotEqual(t, post_1.Title, post_2.Title)
	require.NotEqual(t, post_1.PostCategory, post_2.PostCategory)
	require.NotEqual(t, post_1.Content, post_2.Content)
	require.NotEqual(t, post_1.TimeForRead, post_2.TimeForRead)

}

func TestDeletePost(t *testing.T) {
	category := createRandomCategory(t)
	post_1 := createRandomPost(t, category)
	err := testQueries.DeletePost(context.Background(), post_1.ID)
	require.NoError(t, err)

	post_2, err := testQueries.GetPost(context.Background(), post_1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post_2)
}