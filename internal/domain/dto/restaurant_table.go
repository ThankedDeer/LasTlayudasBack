package dto

import "time"

// RestaurantTableDTO representa la estructura básica de una mesa de restaurante
type RestaurantTableDTO struct {
	TableID   int32     `json:"table_id"`
	Number    int32     `json:"number"`
	StatusID  int32     `json:"status_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateRestaurantTableRequest representa el cuerpo de la solicitud para crear una nueva mesa
type CreateRestaurantTableRequest struct {
	Number   int32 `json:"number" validate:"required"`
	StatusID int32 `json:"status_id" validate:"required"`
}

// UpdateRestaurantTableRequest representa el cuerpo de la solicitud para actualizar una mesa existente
type UpdateRestaurantTableRequest struct {
	TableID  int32 `json:"table_id" validate:"required"`
	Number   int32 `json:"number" validate:"required"`
	StatusID int32 `json:"status_id" validate:"required"`
}

// RestaurantTableResponse representa la respuesta al cliente con los datos de la mesa
type RestaurantTableResponse struct {
	TableID   int32     `json:"table_id"`
	Number    int32     `json:"number"`
	StatusID  int32     `json:"status_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
