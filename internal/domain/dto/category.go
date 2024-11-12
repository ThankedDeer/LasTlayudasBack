package dto

// CreateCategoryRequest representa la solicitud para crear una nueva categoría
type CreateCategoryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
	Is_active   bool    `json:"is_active"`   // Estado activo de la categoría
}

// UpdateCategoryRequest representa la solicitud para actualizar una categoría existente
type UpdateCategoryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
	Is_active   bool    `json:"is_active"`   // Estado activo de la categoría
}
