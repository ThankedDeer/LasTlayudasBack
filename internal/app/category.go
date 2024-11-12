package app

import (
	"context"
	"database/sql"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

// CategoryApp contiene la lógica de negocio para las categorías
type CategoryApp struct {
	store *sqlc.Store
}

// NewCategoryApp crea una nueva instancia de CategoryApp
func NewCategoryApp(store *sqlc.Store) CategoryApp {
	return CategoryApp{
		store: store,
	}
}

// CreateCategory crea una nueva categoría en la base de datos
func (u *CategoryApp) CreateCategory(ctx context.Context, name string, description *string, is_active bool) (sqlc.Category, error) {
	var category sqlc.Category

	err := u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		arg := sqlc.CreateCategoryParams{
			Name: name,
			Description: sql.NullString{
				String: *description,
				Valid:  true,
			},
			IsActive: sql.NullBool{
				Bool:  is_active,
				Valid: false,
			},
		}
		cat, err := q.CreateCategory(ctx, arg)
		if err != nil {
			return err
		}
		category = cat
		return nil
	})

	return category, err
}

// DeleteCategory elimina una categoría de la base de datos por su ID
func (u *CategoryApp) DeleteCategory(ctx context.Context, id int32) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.DeleteCategory(ctx, id)
	})
}

// GetAllCategories obtiene todas las categorías de la base de datos
func (u *CategoryApp) GetAllCategories(ctx context.Context) ([]sqlc.Category, error) {
	return u.store.GetAllCategories(ctx)
}

// GetCategoryByID obtiene una categoría por su ID
func (u *CategoryApp) GetCategoryByID(ctx context.Context, id int32) (sqlc.Category, error) {
	return u.store.GetCategoryByID(ctx, id)
}

// GetCategoryByName obtiene una categoría por su nombre
func (u *CategoryApp) GetCategoryByName(ctx context.Context, name string) (sqlc.Category, error) {
	return u.store.GetCategoryByName(ctx, name)
}

// UpdateCategory actualiza una categoría existente
func (u *CategoryApp) UpdateCategory(ctx context.Context, arg sqlc.UpdateCategoryParams) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.UpdateCategory(ctx, arg)
	})
}
