package app

import (
	"context"
	"database/sql"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

type RolApp struct {
	store *sqlc.Store
}

func NewRoleApp(store *sqlc.Store) RolApp {
	return RolApp{
		store: store,
	}
}

func (u *RolApp) CreateRole(ctx context.Context, name string, description *string) (sqlc.Role, error) {
	var rol sqlc.Role

	err := u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		newRol := sqlc.CreateRoleParams{
			Name: name,
			Description: sql.NullString{
				String: *description,
				Valid:  true,
			},
		}
		cat, err := q.CreateRole(ctx, newRol)
		if err != nil {
			return err
		}
		rol = cat
		return nil
	})

	return rol, err
}

func (u *RolApp) DeleteRole(ctx context.Context, id int32) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.DeleteRole(ctx, id)
	})
}

func (u *CategoryApp) GetAllRoles(ctx context.Context) ([]sqlc.Role, error) {
	return u.store.GetAllRoles(ctx)
}

func (u *CategoryApp) GetRoleByID(ctx context.Context, id int32) (sqlc.Role, error) {
	return u.store.GetRoleByID(ctx, id)
}

func (u *CategoryApp) GetRoleByName(ctx context.Context, name string) (sqlc.Role, error) {
	return u.store.GetRoleByName(ctx, name)
}

func (u *CategoryApp) UpdateRole(ctx context.Context, arg sqlc.UpdateRoleParams) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.UpdateRole(ctx, arg)
	})
}
