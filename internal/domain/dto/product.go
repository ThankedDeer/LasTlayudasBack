package dto

type CreateProductRequest struct {
	Name          string `json:"name"`
	PurchasePrice string `json:"purchase_price"`
	SalePrice     string `json:"sale_price"`
	Stock         int32  `json:"stock"`
	CategoryID    int32  `json:"category_id"`
	ProviderID    int32  `json:"provider_id"`
}

type UpdateProductRequest struct {
	Name          string `json:"name"`
	PurchasePrice string `json:"purchase_price"`
	SalePrice     string `json:"sale_price"`
	Stock         int32  `json:"stock"`
	CategoryID    int32  `json:"category_id"`
	ProviderID    int32  `json:"provider_id"`
}
