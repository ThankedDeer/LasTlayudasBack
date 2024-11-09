package app

import (
	"context"
	"errors"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

type ProviderApp struct {
	store *sqlc.Store
}

func NewProviderApp(store *sqlc.Store) ProviderApp {
	return ProviderApp{
		store: store,
	}
}

func (u *ProviderApp) CreateProvider(data dto.CreateProviderRequest) (sqlc.Provider, error) {
	ctx := context.Background()

	var provider sqlc.Provider
	err := u.store.ExecTx(ctx, func(q *sqlc.Queries) error {
		newProvider := sqlc.CreateProviderParams{
			Name:    data.Name,
			Phone:   data.Phone,
			Email:   data.Email,
			Address: data.Address,
		}

		var err error
		provider, err = q.CreateProvider(ctx, newProvider) // Asignación directa
		return err
	})

	return provider, err
}

func (u *ProviderApp) GetProviders() ([]sqlc.Provider, error) {
	providers, err := u.store.GetAllProviders(context.Background())
	if err != nil {
		return nil, err
	}
	if len(providers) == 0 {
		return nil, errors.New("no se encontraron proveedores")
	}
	return providers, nil
}

func (u *ProviderApp) GetProviderByID(providerID int32) (sqlc.Provider, error) {
	ctx := context.Background()
	provider, err := u.store.GetProviderByID(ctx, providerID)
	if err != nil {
		return sqlc.Provider{}, err
	}
	return provider, nil
}

func (u *ProviderApp) GetProviderByEmail(email string) (sqlc.Provider, error) {
	ctx := context.Background()
	provider, err := u.store.GetProviderByEmail(ctx, email)
	if err != nil {
		return sqlc.Provider{}, err
	}
	return provider, nil
}

func (u *ProviderApp) UpdateProvider(data dto.UpdateProviderRequest) error {
	ctx := context.Background()

	// Convierte dto.UpdateProviderRequest a sqlc.UpdateProviderParams
	updateParams := sqlc.UpdateProviderParams{
		ProviderID: data.ProviderID,
		Name:       data.Name,
		Phone:      data.Phone,
		Email:      data.Email,
		Address:    data.Address,
	}

	return u.store.UpdateProvider(ctx, updateParams)
}

func (u *ProviderApp) DeleteProvider(providerID int32) error {
	ctx := context.Background()
	return u.store.DeleteProvider(ctx, providerID)
}
