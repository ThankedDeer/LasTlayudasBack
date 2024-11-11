package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"

	"github.com/labstack/echo/v4"
)

type IProviderHandler interface {
	CreateProvider(c echo.Context) error
	GetAllProvider(c echo.Context) error
	UpdateProvider(c echo.Context) error
	GetProviderByIDHandler(c echo.Context) error
	GetProviderByEmailHandler(c echo.Context) error
	DeleteProviderHandler(c echo.Context) error
}

type ProviderHandler struct {
	app app.ProviderApp
}

func NewProviderHandler(app app.ProviderApp) IProviderHandler {
	return &ProviderHandler{
		app: app,
	}
}

// CreateProvider maneja la creación de un proveedor
func (u *ProviderHandler) CreateProvider(c echo.Context) error {
	var req dto.CreateProviderRequest

	// Decodifica el JSON usando json.NewDecoder para un manejo más preciso de los errores
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		log.Printf("Error en el Bind: %v", err)
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
	}

	log.Printf("Datos recibidos: %+v\n", req) // Verifica los datos recibidos

	provider, err := u.app.CreateProvider(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo crear el proveedor"})
	}

	return c.JSON(http.StatusCreated, provider)
}

// GetAllProvider maneja la obtención de todos los proveedores
func (u *ProviderHandler) GetAllProvider(c echo.Context) error {
	providers, err := u.app.GetProviders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, providers)
}

// GetProviderByIDHandler maneja la obtención de un proveedor por ID
func (u *ProviderHandler) GetProviderByIDHandler(c echo.Context) error {
	idParam := c.Param("id")
	providerID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	provider, err := u.app.GetProviderByID(int32(providerID))
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Proveedor no encontrado"})
	}

	return c.JSON(http.StatusOK, provider)
}

// GetProviderByEmailHandler maneja la obtención de un proveedor por correo electrónico
func (u *ProviderHandler) GetProviderByEmailHandler(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Correo electrónico es requerido"})
	}

	provider, err := u.app.GetProviderByEmail(email)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "Proveedor no encontrado"})
	}

	return c.JSON(http.StatusOK, provider)
}

// UpdateProvider maneja la actualización de un proveedor
func (u *ProviderHandler) UpdateProvider(c echo.Context) error {
	var req dto.UpdateProviderRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Datos inválidos"})
	}

	err := u.app.UpdateProvider(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo actualizar el proveedor"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Proveedor actualizado exitosamente"})
}

// DeleteProviderHandler maneja la eliminación de un proveedor por ID
func (u *ProviderHandler) DeleteProviderHandler(c echo.Context) error {
	idParam := c.Param("id")
	providerID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "ID inválido"})
	}

	err = u.app.DeleteProvider(int32(providerID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "No se pudo eliminar el proveedor"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Proveedor eliminado exitosamente"})
}
