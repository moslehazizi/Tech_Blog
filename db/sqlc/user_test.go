package db

import (
	"Tech_Blog/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       util.RandomName(),
		HashedPassword: util.RandomName(),
		FullName:       util.RandomName(),
		Email:          util.RandomName(),
		UserImage:      util.RandomBytes(10),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.UserImage, user.UserImage)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user_1 := createRandomUser(t)
	user_2, err := testQueries.GetUser(context.Background(), user_1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user_2)
	require.Equal(t, user_1.Username, user_2.Username)
	require.Equal(t, user_1.HashedPassword, user_2.HashedPassword)
	require.Equal(t, user_1.FullName, user_2.FullName)
	require.Equal(t, user_1.Email, user_2.Email)
	require.Equal(t, user_1.UserImage, user_2.UserImage)
}

func TestUpdateUser(t *testing.T) {
	user_1 := createRandomUser(t)

	arg := UpdateUserParams{
		Username:       util.RandomName(),
		FullName:       util.RandomName(),
		HashedPassword: util.RandomName(),
		Email:          user_1.Email,
		UserImage:      util.RandomBytes(20),
	}
	user_2, err := testQueries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user_2)
	require.Equal(t, user_1.Email, user_2.Email)
	require.NotEqual(t, user_1.HashedPassword, user_2.HashedPassword)
	require.NotEqual(t, user_1.FullName, user_2.FullName)
	require.NotEqual(t, user_1.Username, user_2.Username)
	require.NotEqual(t, user_1.UserImage, user_2.UserImage)

}

func TestDeleteUser(t *testing.T) {
	user_1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user_1.Username)

	require.NoError(t, err)
	
	user_2, err := testQueries.GetUser(context.Background(), user_1.Username)
	
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user_2)
}