package app

import (
	"context"
	"errors"

	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

type ProductApp struct {
	store *sqlc.Store
}

func NewProductApp(store *sqlc.Store) ProductApp {
	return ProductApp{
		store: store,
	}
}

func (u *ProductApp) CreateProduct(data dto.CreateProductRequest) error {
	ctx := context.Background()

	err := u.store.ExecTx(ctx, func(u *sqlc.Queries) error {
		newProduct := sqlc.CreateProductParams{
			Name:          data.Name,
			PurchasePrice: data.PurchasePrice,
			SalePrice:     data.SalePrice,
			Stock:         data.Stock,
			CategoryID:    data.CategoryID,
			ProviderID:    data.ProviderID,
		}
		_, err := u.CreateProduct(ctx, newProduct)
		return err
	})
	return err
}

func (u *ProductApp) GetProduct() ([]sqlc.Product, error) {
	products, err := u.store.GetAllProducts(context.Background())
	if err != nil {
		return nil, err
	}
	if len(products) == 0 {
		return nil, errors.New("no se encontraron productos")
	}
	return products, nil
}

func (u *ProductApp) UpdateProduct(id int, data dto.UpdateProductRequest) error {
	arg := sqlc.UpdateProductParams{
		Name:          data.Name,
		PurchasePrice: data.PurchasePrice,
		SalePrice:     data.SalePrice,
		Stock:         data.Stock,
		CategoryID:    data.CategoryID,
		ProviderID:    data.ProviderID,
	}
	_, err := u.store.UpdateProduct(context.Background(), arg)
	return err
}
