package db

import (
	"context"
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

func CreatRandomeUser(t *testing.T) Users {
	f := faker.New()

	arg := CreateUserParams{
		Firstname: f.Person().FirstName(),
		Lastname:  f.Person().LastName(),
		Password:  f.Internet().Password(),
		Email:     f.UUID().V4() + f.Internet().Email(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Firstname, user.Firstname)
	require.Equal(t, arg.Lastname, user.Lastname)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	CreatRandomeUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := CreatRandomeUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Email)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.Firstname, user2.Firstname)
	require.Equal(t, user1.Lastname, user2.Lastname)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}

func TestUpdatUserPassword(t *testing.T) {
	f := faker.New()
	user1 := CreatRandomeUser(t)

	arg := UpdatPasswordParams{
		Email:    user1.Email,
		Password: f.Internet().Password(),
	}

	user2, err := testQueries.UpdatPassword(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.Firstname, user2.Firstname)
	require.Equal(t, user1.Lastname, user2.Lastname)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, arg.Password, user2.Password)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)

}
