package handler

import (
	"encoding/json"
	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

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
// @Summary Create a new product
// @Description Create a new product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Product to create"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/products/ [post]
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
// @Summary Get all products
// @Description Retrieve all products
// @Tags products
// @Produce json
// @Failure 404 {object} map[string]string
// @Router /api/products/ [get]
func (u *ProductHandler) GetAllProducts(c echo.Context) error {
	products, err := u.app.GetProduct()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Productos no encontrados"})
	}

	return c.JSON(http.StatusOK, products)
}

// UpdateProduct handles product updates
// @Summary Update an existing product
// @Description Update an existing product with the input payload
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/products/{id} [put]
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
