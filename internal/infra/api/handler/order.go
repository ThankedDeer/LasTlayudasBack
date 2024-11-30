package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

type IOrderHandler interface {
	CreateOrder(c echo.Context) error
	GetAllOrders(c echo.Context) error
	UpdateOrder(c echo.Context) error
	DeleteOrder(c echo.Context) error
	GetOrderByID(c echo.Context) error
}

type OrderHandler struct {
	app app.OrderApp
}

func NewOrderHandler(app app.OrderApp) IOrderHandler {
	return &OrderHandler{
		app: app,
	}
}

// CreateOrder maneja la creación de una nueva orden
func (u *OrderHandler) CreateOrder(c echo.Context) error {
	var req dto.CreateOrderRequest

	// Decodificar el cuerpo JSON en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos: " + err.Error()})
	}

	order, err := u.app.CreateOrder(c.Request().Context(), req.TableID, req.StatusID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, order)
}

// GetAllOrders obtiene todas las órdenes
func (u *OrderHandler) GetAllOrders(c echo.Context) error {
	orders, err := u.app.GetAllOrders(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Órdenes no encontradas"})
	}

	return c.JSON(http.StatusOK, orders)
}

// UpdateOrder maneja la actualización de una orden
func (u *OrderHandler) UpdateOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de orden inválido"})
	}

	var req dto.UpdateOrderRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	updateParams := sqlc.UpdateOrderParams{
		OrderID:  int32(id),
		TableID:  req.TableID,
		StatusID: req.StatusID,
	}

	if err := u.app.UpdateOrder(c.Request().Context(), updateParams); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Orden actualizada con éxito"})
}

// DeleteOrder maneja la eliminación de una orden por su ID
func (u *OrderHandler) DeleteOrder(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de orden inválido"})
	}

	if err := u.app.DeleteOrder(c.Request().Context(), int32(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Orden eliminada correctamente"})
}

// GetOrderByID obtiene una orden por su ID
func (u *OrderHandler) GetOrderByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de orden inválido"})
	}

	order, err := u.app.GetOrderByID(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Orden no encontrada"})
	}

	return c.JSON(http.StatusOK, order)
}
