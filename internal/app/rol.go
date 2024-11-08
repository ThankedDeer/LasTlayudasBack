package app

import (
	"context"
	"database/sql"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

type RoleApp struct {
	store *sqlc.Store
}

func NewRoleApp(store *sqlc.Store) RoleApp {
	return RoleApp{
		store: store,
	}
}

func (u *RoleApp) CreateRole(ctx context.Context, name string, description *string) (sqlc.Role, error) {
	var role sqlc.Role

	err := u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		newRol := sqlc.CreateRoleParams{
			Name: name,
			Description: sql.NullString{
				String: *description,
				Valid:  true,
			},
		}
		rol, err := q.CreateRole(ctx, newRol)
		if err != nil {
			return err
		}
		role = rol
		return nil
	})

	return role, err
}

func (u *RoleApp) DeleteRole(ctx context.Context, id int32) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.DeleteRole(ctx, id)
	})
}

func (u *RoleApp) GetAllRoles(ctx context.Context) ([]sqlc.Role, error) {
	return u.store.GetAllRoles(ctx)
}

func (u *RoleApp) GetRoleByID(ctx context.Context, id int32) (sqlc.Role, error) {
	return u.store.GetRoleByID(ctx, id)
}

func (u *RoleApp) GetRoleByName(ctx context.Context, name string) (sqlc.Role, error) {
	return u.store.GetRoleByName(ctx, name)
}

func (u *RoleApp) UpdateRole(ctx context.Context, arg sqlc.UpdateRoleParams) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.UpdateRole(ctx, arg)
	})
}
