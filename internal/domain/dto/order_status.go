package dto

type CreateOrderStatusRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
}

type UpdateOrderStatusRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
}
