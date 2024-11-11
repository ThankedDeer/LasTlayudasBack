package app

import (
	"context"
	"database/sql"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

type OrderStatusApp struct {
	store *sqlc.Store
}

func NewOrderStatusApp(store *sqlc.Store) OrderStatusApp {
	return OrderStatusApp{
		store: store,
	}
}

func (u *OrderStatusApp) CreateOrderStatus(ctx context.Context, name string, description *string) (sqlc.OrderStatus, error) {
	var orderStatus sqlc.OrderStatus

	err := u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		newOrderStatus := sqlc.CreateOrderStatusParams{
			Name: name,
			Description: sql.NullString{
				String: *description,
				Valid:  true,
			},
		}
		ord_status, err := q.CreateOrderStatus(ctx, newOrderStatus)
		if err != nil {
			return err
		}
		orderStatus = ord_status
		return nil
	})

	return orderStatus, err
}

func (u *OrderStatusApp) DeleteOrderStatus(ctx context.Context, id int32) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.DeleteOrderStatus(ctx, id)
	})
}

func (u *OrderStatusApp) GetAllOrderStatus(ctx context.Context) ([]sqlc.OrderStatus, error) {
	return u.store.GetAllOrderStatuses(ctx)
}

func (u *OrderStatusApp) GetOrderStatusByID(ctx context.Context, id int32) (sqlc.OrderStatus, error) {
	return u.store.GetOrderStatusByID(ctx, id)
}

func (u *OrderStatusApp) UpdateOrderStatus(ctx context.Context, arg sqlc.UpdateOrderStatusParams) error {
	return u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		return q.UpdateOrderStatus(ctx, arg)
	})
}
