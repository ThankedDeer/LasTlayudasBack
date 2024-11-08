package dto

// CreateCategoryRequest representa la solicitud para crear una nueva categoría
type CreateRoleRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
}

// UpdateCategoryRequest representa la solicitud para actualizar una categoría existente
type UpdateRoleRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"` // Permite descripciones nulas
}
