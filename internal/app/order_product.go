package app

import (
	"context"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

// OrderProductApp contiene la lógica de negocio para order_product
type OrderProductApp struct {
	store *sqlc.Store
}

// NewOrderProductApp crea una nueva instancia de OrderProductApp
func NewOrderProductApp(store *sqlc.Store) OrderProductApp {
	return OrderProductApp{
		store: store,
	}
}

// CreateOrderProduct crea un nuevo registro en la tabla order_product
func (u *OrderProductApp) CreateOrderProduct(ctx context.Context, orderID int32, productID int32, quantity int32) (sqlc.OrderProduct, error) {
	var orderProduct sqlc.OrderProduct

	err := u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		arg := sqlc.CreateOrderProductParams{
			OrderID:   orderID,
			ProductID: productID,
			Column3:   quantity,
		}
		op, err := q.CreateOrderProduct(ctx, arg)
		if err != nil {
			return err
		}
		orderProduct = op
		return nil
	})

	return orderProduct, err
}

// DeleteOrderProduct elimina un registro de order_product por su ID
func (u *OrderProductApp) DeleteOrderProduct(ctx context.Context, id int32) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.DeleteOrderProduct(ctx, id)
	})
}

// GetOrderProductByID obtiene un registro de order_product por su ID
func (u *OrderProductApp) GetOrderProductByID(ctx context.Context, id int32) (sqlc.OrderProduct, error) {
	return u.store.GetOrderProductByID(ctx, id)
}

// UpdateOrderProduct actualiza un registro de order_product existente
func (u *OrderProductApp) UpdateOrderProduct(ctx context.Context, arg sqlc.UpdateOrderProductParams) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.UpdateOrderProduct(ctx, arg)
	})
}
