package dto

type CreateOrderProductRequest struct {
	OrderID   int32 `json:"order_id" validate:"required"`
	ProductID int32 `json:"product_id" validate:"required"`
	Quantity  int32 `json:"quantity" validate:"min=1"`
}

type UpdateOrderProductRequest struct {
	OrderProductID int32 `json:"order_product_id" validate:"required"`
	Quantity       int32 `json:"quantity" validate:"min=1"`
}
