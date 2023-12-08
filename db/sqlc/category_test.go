package db

import (
	"Tech_Blog/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomCategory(t *testing.T) Category {
	name_of_category := util.RandomName()
	category, err := testQueries.CreateCategory(context.Background(), name_of_category)

	require.NoError(t, err)
	require.NotEmpty(t, category)
	require.Equal(t, name_of_category, category.CategoryName)
	require.NotZero(t, category.ID)

	return category
}

func TestCreateCategory(t *testing.T) {
	createRandomCategory(t)
}

func TestGetCategory(t *testing.T) {
	category_1 := createRandomCategory(t)
	category_2, err := testQueries.GetCategory(context.Background(), category_1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, category_2)
	require.Equal(t, category_1.ID, category_2.ID)
	require.Equal(t, category_1.CategoryName, category_2.CategoryName)
}

func TestUpdateCategory(t *testing.T) {
	category_1 := createRandomCategory(t)
	arg := UpdateCategoryParams{
		ID:           category_1.ID,
		CategoryName: util.RandomName(),
	}
	category_2, err := testQueries.UpdateCategory(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, category_2)
	require.Equal(t, category_1.ID, category_2.ID)
	require.NotEqual(t, category_1.CategoryName, category_2.CategoryName)
}

func TestDeleteCategory(t *testing.T) {
	category_1 := createRandomCategory(t)
	err := testQueries.DeleteCategory(context.Background(), category_1.ID)
	require.NoError(t, err)

	category_2, err := testQueries.GetCategory(context.Background(), category_1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, category_2)

}
