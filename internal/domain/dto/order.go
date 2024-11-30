package dto

type CreateOrderRequest struct {
	TableID  int32 `json:"table_id"`  // ID de la mesa asociada con la orden
	StatusID int32 `json:"status_id"` // ID del estado inicial de la orden
}

type UpdateOrderRequest struct {
	TableID  int32 `json:"table_id"`  // Nuevo ID de la mesa (opcional si no cambia)
	StatusID int32 `json:"status_id"` // Nuevo estado de la orden
}

type OrderResponse struct {
	OrderID   int32  `json:"order_id"`   // ID único de la orden
	TableID   int32  `json:"table_id"`   // ID de la mesa asociada
	StatusID  int32  `json:"status_id"`  // ID del estado de la orden
	CreatedAt string `json:"created_at"` // Fecha de creación
	UpdatedAt string `json:"updated_at"` // Última f
}
