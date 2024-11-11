package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
)

type IProductHandler interface {
	CreateProduct(c echo.Context) error
	GetAllProducts(c echo.Context) error
	UpdateProduct(c echo.Context) error
}

type ProductHandler struct {
	app app.ProductApp
}

func NewProductHandler(app app.ProductApp) IProductHandler {
	return &ProductHandler{
		app: app,
	}
}

// CreateProduct handles product creation
func (u *ProductHandler) CreateProduct(c echo.Context) error {
	var req dto.CreateProductRequest

	// Usar NewDecoder para decodificar el cuerpo JSON en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Llamar al método de la aplicación para crear el producto
	if err := u.app.CreateProduct(req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Producto creado con éxito"})
}

// GetAllProducts retrieves all products
func (u *ProductHandler) GetAllProducts(c echo.Context) error {
	products, err := u.app.GetProduct()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Productos no encontrados"})
	}

	return c.JSON(http.StatusOK, products)
}

// UpdateProduct handles product updates
func (u *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product ID"})
	}

	var req dto.UpdateProductRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if err := u.app.UpdateProduct(id, req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Producto actualizado con éxito"})
}
