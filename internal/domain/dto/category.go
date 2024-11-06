package dto

// CreateCategoryRequest representa la solicitud para crear una nueva categoría
type CreateCategoryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
	Column3     bool    `json:"active"`      // Estado activo de la categoría
}

// UpdateCategoryRequest representa la solicitud para actualizar una categoría existente
type UpdateCategoryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
	Active      *bool   `json:"active"`      // Estado activo de la categoría
}
