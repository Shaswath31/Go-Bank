package db

import (
	"context"
	"testing"

	"github.com/shaswath/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	hashedPassword, err := utils.HashPassword(utils.RandomString(6))

	require.NoError(t, err)
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}

	user, err := testStore.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	gotUser, err := testStore.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.Equal(t, user, gotUser)
}
