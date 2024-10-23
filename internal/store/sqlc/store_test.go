package sqlc

import (
	"context"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/require"
)

func TestUpdateUserPasswordTx(t *testing.T) {
	f := faker.New()
	store := NewStore(testDB)

	user := CreatRandomeUser(t)

	n := 5
	errs := make(chan error)
	results := make(chan updatedPasswordTxResult)
	for i := 0; i < n; i++ {
		go func() {
			result, err := store.UpdateUserPasswordTx(context.Background(), UpdatePasswordTxParams{
				Email:    user.Email,
				Password: f.Internet().Password(),
			})

			errs <- err
			results <- result

		}()

	}
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

	}

}
