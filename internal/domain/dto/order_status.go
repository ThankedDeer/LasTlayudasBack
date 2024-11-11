package dto

// CreateCategoryRequest representa la solicitud para crear una nueva categoría
type CreateOrderStatusRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
}

// UpdateCategoryRequest representa la solicitud para actualizar una categoría existente
type UpdateOrderStatusRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
}
