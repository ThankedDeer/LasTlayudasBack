package app

import (
	"context"
	"errors"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

// RestaurantTableApp contiene la lógica de negocio para las mesas de restaurante
type RestaurantTableApp struct {
	store *sqlc.Store
}

// NewRestaurantTableApp crea una nueva instancia de RestaurantTableApp
func NewRestaurantTableApp(store *sqlc.Store) RestaurantTableApp {
	return RestaurantTableApp{
		store: store,
	}
}

// CreateRestaurantTable crea una nueva mesa de restaurante en la base de datos
func (u *RestaurantTableApp) CreateRestaurantTable(ctx context.Context, number int32, statusID int32) (sqlc.RestaurantTable, error) {
	var table sqlc.RestaurantTable

	err := u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		arg := sqlc.CreateRestaurantTableParams{
			Number:   number,
			StatusID: statusID,
		}
		t, err := q.CreateRestaurantTable(ctx, arg)
		if err != nil {
			return err
		}
		table = t
		return nil
	})

	return table, err
}

// DeleteRestaurantTable elimina una mesa de restaurante de la base de datos por su ID
func (u *RestaurantTableApp) DeleteRestaurantTable(ctx context.Context, tableID int32) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.DeleteRestaurantTable(ctx, tableID)
	})
}

// GetAllRestaurantTables obtiene todas las mesas de restaurante de la base de datos
func (u *RestaurantTableApp) GetAllRestaurantTables(ctx context.Context) ([]sqlc.RestaurantTable, error) {
	tables, err := u.store.GetAllRestaurantTables(ctx)
	if err != nil {
		return nil, err
	}
	if len(tables) == 0 {
		return nil, errors.New("no se encontraron mesas")
	}
	return tables, nil
}

// GetRestaurantTableByID obtiene una mesa de restaurante por su ID
func (u *RestaurantTableApp) GetRestaurantTableByID(ctx context.Context, tableID int32) (sqlc.RestaurantTable, error) {
	return u.store.GetRestaurantTableByID(ctx, tableID)
}

// UpdateRestaurantTable actualiza una mesa existente
func (u *RestaurantTableApp) UpdateRestaurantTable(ctx context.Context, tableID int32, number int32, statusID int32) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		arg := sqlc.UpdateRestaurantTableParams{
			TableID:  tableID,
			Number:   number,
			StatusID: statusID,
		}
		return q.UpdateRestaurantTable(ctx, arg)
	})
}
