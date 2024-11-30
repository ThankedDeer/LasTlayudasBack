package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

// IOrderProductHandler define las operaciones para manejar order_product
type IOrderProductHandler interface {
	CreateOrderProduct(c echo.Context) error
	UpdateOrderProduct(c echo.Context) error
	GetOrderProductByID(c echo.Context) error
	DeleteOrderProduct(c echo.Context) error
}

// OrderProductHandler implementa la interfaz IOrderProductHandler
type OrderProductHandler struct {
	app app.OrderProductApp
}

// NewOrderProductHandler crea una nueva instancia de OrderProductHandler
func NewOrderProductHandler(app app.OrderProductApp) IOrderProductHandler {
	return &OrderProductHandler{
		app: app,
	}
}

// GetOrderProductByID obtiene un registro de order_product por su ID
func (h *OrderProductHandler) GetOrderProductByID(c echo.Context) error {
	// Obtener el ID del parámetro de la ruta
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order_product ID"})
	}

	// Llamar al método de la aplicación para obtener el registro por ID
	orderProduct, err := h.app.GetOrderProductByID(c.Request().Context(), int32(id))
	if err != nil {
		// Si no se encuentra el registro, retornar un error
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "OrderProduct not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retornar el registro encontrado
	return c.JSON(http.StatusOK, orderProduct)
}

// CreateOrderProduct maneja la creación de registros en order_product
func (h *OrderProductHandler) CreateOrderProduct(c echo.Context) error {
	var req dto.CreateOrderProductRequest

	// Usar NewDecoder para decodificar el cuerpo JSON en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Llamar al método de la aplicación para crear el registro
	orderProduct, err := h.app.CreateOrderProduct(c.Request().Context(), req.OrderID, req.ProductID, req.Quantity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, orderProduct)
}

// DeleteOrderProduct elimina un registro de order_product por su ID
func (h *OrderProductHandler) DeleteOrderProduct(c echo.Context) error {
	// Obtener el ID del parámetro de la ruta
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order_product ID"})
	}

	// Llamar al método de la aplicación para eliminar el registro
	err = h.app.DeleteOrderProduct(c.Request().Context(), int32(id))
	if err != nil {
		// Si no se encuentra el registro o hay otro error, manejar el error apropiadamente
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "OrderProduct not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Responder con un mensaje de éxito
	return c.JSON(http.StatusOK, map[string]string{"message": "Registro eliminado con éxito"})
}

// UpdateOrderProduct maneja la actualización de registros en order_product
func (h *OrderProductHandler) UpdateOrderProduct(c echo.Context) error {
	// Obtener el ID del parámetro de la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid order_product ID"})
	}

	var req dto.UpdateOrderProductRequest
	// Decodificar el cuerpo de la solicitud en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Crear los parámetros para actualizar
	arg := sqlc.UpdateOrderProductParams{
		OrderProductID: int32(id), // Usar el ID convertido de la URL
		Quantity:       req.Quantity,
	}

	// Llamar al método de la aplicación para actualizar el registro
	if err := h.app.UpdateOrderProduct(c.Request().Context(), arg); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Registro actualizado con éxito"})
}
