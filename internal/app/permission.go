package app

import (
	"context"
	"database/sql"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

// PermissionApp contiene la lógica de negocio para los permisos
type PermissionApp struct {
	store *sqlc.Store
}

// NewPermissionApp crea una nueva instancia de PermissionApp
func NewPermissionApp(store *sqlc.Store) PermissionApp {
	return PermissionApp{
		store: store,
	}
}

// CreatePermission crea un nuevo permiso en la base de datos
func (p *PermissionApp) CreatePermission(ctx context.Context, name string, description *string) (sqlc.Permission, error) {
	var permission sqlc.Permission

	err := p.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		arg := sqlc.CreatePermissionParams{
			Name: name,
			Description: sql.NullString{
				String: *description,
				Valid:  description != nil,
			},
		}
		perm, err := q.CreatePermission(ctx, arg)
		if err != nil {
			return err
		}
		permission = perm
		return nil
	})

	return permission, err
}

// DeletePermission elimina un permiso de la base de datos por su ID
func (p *PermissionApp) DeletePermission(ctx context.Context, id int32) error {
	return p.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.DeletePermission(ctx, id)
	})
}

// GetAllPermissions obtiene todos los permisos de la base de datos
func (p *PermissionApp) GetAllPermissions(ctx context.Context) ([]sqlc.Permission, error) {
	return p.store.GetAllPermissions(ctx)
}

// GetPermissionByID obtiene un permiso por su ID
func (p *PermissionApp) GetPermissionByID(ctx context.Context, id int32) (sqlc.Permission, error) {
	return p.store.GetPermissionByID(ctx, id)
}

// UpdatePermission actualiza un permiso existente
func (p *PermissionApp) UpdatePermission(ctx context.Context, arg sqlc.UpdatePermissionParams) error {
	return p.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.UpdatePermission(ctx, arg)
	})
}
