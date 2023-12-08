package db

import (
	"Tech_Blog/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomReview(t *testing.T, user User, post Post) Review {
	arg := CreateReviewParams{
		Reviewer:      user.Username,
		ReviewContent: util.RandomName(),
		Post:          post.ID,
		StarDegree:    util.RandomFloat(),
	}
	review, err := testQueries.CreateReview(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, review)
	require.Equal(t, arg.Reviewer, review.Reviewer)
	require.Equal(t, arg.ReviewContent, review.ReviewContent)
	require.Equal(t, arg.Post, review.Post)
	require.Equal(t, arg.StarDegree, review.StarDegree)
	require.NotZero(t, review.ID)

	return review
}

func TestCreateReview(t *testing.T) {
	category := createRandomCategory(t)
	user := createRandomUser(t)
	post := createRandomPost(t, category)
	createRandomReview(t, user, post)
}

func TestGetReview(t *testing.T) {
	category := createRandomCategory(t)
	user := createRandomUser(t)
	post := createRandomPost(t, category)
	review_1 := createRandomReview(t, user, post)
	review_2, err := testQueries.GetReview(context.Background(), review_1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, review_2)
	require.Equal(t, review_1.ID, review_2.ID)
	require.Equal(t, review_1.Reviewer, review_2.Reviewer)
	require.Equal(t, review_1.ReviewContent, review_2.ReviewContent)
	require.Equal(t, review_1.Post, review_2.Post)
	require.Equal(t, review_1.StarDegree, review_2.StarDegree)

}

func TestDeleteRrview(t *testing.T) {
	category := createRandomCategory(t)
	user := createRandomUser(t)
	post := createRandomPost(t, category)
	review_1 := createRandomReview(t, user, post)

	err := testQueries.DeleteReview(context.Background(), review_1.ID)

	require.NoError(t, err)

	review_2, err := testQueries.GetReview(context.Background(), review_1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, review_2)

}


func TestListReviews(t *testing.T) {
	category := createRandomCategory(t)
	user := createRandomUser(t)
	post := createRandomPost(t, category)
	for i:=0; i<10; i++ {
		createRandomReview(t, user, post)
	}
	arg := ListReviewsParams{
		Post: post.ID,
		Limit: 5,
		Offset: 5,
	}

	reviews, err := testQueries.ListReviews(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, reviews, 5)
	
	for _, review := range reviews {
		require.NotEmpty(t, review)
		require.Equal(t, post.ID, review.Post)
		require.Equal(t, user.Username, review.Reviewer)
	}
}