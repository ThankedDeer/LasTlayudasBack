package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

type IOrderStatusHandler interface {
	CreateOrderStatus(c echo.Context) error
	GetAllOrderStatus(c echo.Context) error
	UpdateOrderStatus(c echo.Context) error
	DeleteOrderStatus(c echo.Context) error
	GetOrderStatusByID(c echo.Context) error
}

type OrderStatusHandler struct {
	app app.OrderStatusApp
}

func NewOrderStatusHandler(app app.OrderStatusApp) IOrderStatusHandler {
	return &OrderStatusHandler{
		app: app,
	}
}

func (u *OrderStatusHandler) CreateOrderStatus(c echo.Context) error {
	var req dto.CreateOrderStatusRequest

	// Decodificar el cuerpo JSON en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos: " + err.Error()})
	}

	// Validación: asegurarse de que 'Name' no esté vacío
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'nombre' es obligatorio."})
	}

	role, err := u.app.CreateOrderStatus(c.Request().Context(), req.Name, req.Description)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, role)
}

func (u *OrderStatusHandler) GetAllOrderStatus(c echo.Context) error {
	roles, err := u.app.GetAllOrderStatus(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Estatus de ordenes no encontrados"})
	}

	return c.JSON(http.StatusOK, roles)
}

func (u *OrderStatusHandler) UpdateOrderStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID del estatus inválido"})
	}

	var req dto.UpdateOrderStatusRequest
	// Decodificar el JSON del cuerpo de la solicitud en req
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		log.Println("Error de decodificación:", err) // Log para el error de decodificación
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	// Convertir *string a sql.NullString para el campo Description
	description := sql.NullString{}
	if req.Description != nil {
		description = sql.NullString{String: *req.Description, Valid: true}
	}

	// Crear los parámetros de actualización usando los datos decodificados
	updateParams := sqlc.UpdateOrderStatusParams{
		OrderStatusID: int32(id),
		Name:          req.Name,
		Description:   description,
	}

	if err := u.app.UpdateOrderStatus(c.Request().Context(), updateParams); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Respuesta de éxito si la actualización fue exitosa
	return c.JSON(http.StatusOK, map[string]string{"message": "Estatus actualizado con éxito"})
}

func (u *OrderStatusHandler) DeleteOrderStatus(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID del estatus inválido"})
	}

	if err := u.app.DeleteOrderStatus(c.Request().Context(), int32(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Responder con un mensaje de éxito y código 200 OK
	return c.JSON(http.StatusOK, map[string]string{"message": "Estatus eliminado correctamente"})
}

func (u *OrderStatusHandler) GetOrderStatusByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID del estatus inválido"})
	}

	orderStatus, err := u.app.GetOrderStatusByID(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Estatus no encontrado"})
	}

	return c.JSON(http.StatusOK, orderStatus)
}
