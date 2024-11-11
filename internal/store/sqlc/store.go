package sqlc

import (
	"context"
	"database/sql"
	"fmt"

)

// Store provides all function to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx exectes a function within a database transaction
func (Store *Store) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := Store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}
	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()

}

//type UpdatePasswordTxParams struct {
//Email    string `json:"email"`
//Password string `json:"password"`
//}
//type updatedPasswordTxResult struct {
//User Users `json:"user"`
//}

//func (store *Store) UpdateUserPasswordTx(ctx context.Context, arg UpdatePasswordTxParams) (updatedPasswordTxResult, error) {

//var result updatedPasswordTxResult

//err := store.execTx(ctx, func(q *Queries) error {
//var err error

//result.User, err = q.GetUserForUpdate(ctx, 1)

//if err != nil {
//return err
//}

//result.User, err = q.UpdatPassword(ctx, UpdatPasswordParams(arg))

//if err != nil {
//return err
//}

//return nil

//
//)

//return result, err

//}
