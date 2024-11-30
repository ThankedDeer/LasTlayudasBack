package app

import (
	"context"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

// OrderApp contiene la lógica de negocio para las órdenes
type OrderApp struct {
	store *sqlc.Store
}

// NewOrderApp crea una nueva instancia de OrderApp
func NewOrderApp(store *sqlc.Store) OrderApp {
	return OrderApp{
		store: store,
	}
}

// CreateOrder crea una nueva orden en la base de datos
func (u *OrderApp) CreateOrder(ctx context.Context, tableID, statusID int32) (sqlc.Order, error) {
	var order sqlc.Order

	err := u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		arg := sqlc.CreateOrderParams{
			TableID:  tableID,
			StatusID: statusID,
		}
		ord, err := q.CreateOrder(ctx, arg)
		if err != nil {
			return err
		}
		order = ord
		return nil
	})

	return order, err
}

// DeleteOrder elimina una orden de la base de datos por su ID
func (u *OrderApp) DeleteOrder(ctx context.Context, id int32) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.DeleteOrder(ctx, id)
	})
}

// GetAllOrders obtiene todas las órdenes de la base de datos
func (u *OrderApp) GetAllOrders(ctx context.Context) ([]sqlc.Order, error) {
	return u.store.GetAllOrders(ctx)
}

// GetOrderByID obtiene una orden por su ID
func (u *OrderApp) GetOrderByID(ctx context.Context, id int32) (sqlc.Order, error) {
	return u.store.GetOrderByID(ctx, id)
}

// UpdateOrder actualiza una orden existente
func (u *OrderApp) UpdateOrder(ctx context.Context, arg sqlc.UpdateOrderParams) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.UpdateOrder(ctx, arg)
	})
}
