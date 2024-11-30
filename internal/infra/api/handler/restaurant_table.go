package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
)

// IRestaurantTableHandler define las operaciones para manejar restaurant_table
type IRestaurantTableHandler interface {
	CreateRestaurantTable(c echo.Context) error
	UpdateRestaurantTable(c echo.Context) error
	GetRestaurantTableByID(c echo.Context) error
	DeleteRestaurantTable(c echo.Context) error
	GetAllRestaurantTables(c echo.Context) error
}

// RestaurantTableHandler implementa la interfaz IRestaurantTableHandler
type RestaurantTableHandler struct {
	app app.RestaurantTableApp
}

// NewRestaurantTableHandler crea una nueva instancia de RestaurantTableHandler
func NewRestaurantTableHandler(app app.RestaurantTableApp) IRestaurantTableHandler {
	return &RestaurantTableHandler{
		app: app,
	}
}

// GetAllRestaurantTables obtiene todas las mesas de restaurante
func (h *RestaurantTableHandler) GetAllRestaurantTables(c echo.Context) error {
	// Llamar al método de la aplicación para obtener todas las mesas
	tables, err := h.app.GetAllRestaurantTables(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retornar las mesas encontradas
	return c.JSON(http.StatusOK, tables)
}

// GetRestaurantTableByID obtiene un registro de restaurant_table por su ID
func (h *RestaurantTableHandler) GetRestaurantTableByID(c echo.Context) error {
	// Obtener el ID del parámetro de la ruta
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid restaurant_table ID"})
	}

	// Llamar al método de la aplicación para obtener el registro por ID
	table, err := h.app.GetRestaurantTableByID(c.Request().Context(), int32(id))
	if err != nil {
		// Si no se encuentra el registro, retornar un error
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "RestaurantTable not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Retornar el registro encontrado
	return c.JSON(http.StatusOK, table)
}

// CreateRestaurantTable maneja la creación de registros en restaurant_table
func (h *RestaurantTableHandler) CreateRestaurantTable(c echo.Context) error {
	var req dto.CreateRestaurantTableRequest

	// Usar NewDecoder para decodificar el cuerpo JSON en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Llamar al método de la aplicación para crear el registro
	table, err := h.app.CreateRestaurantTable(c.Request().Context(), req.Number, req.StatusID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, table)
}

// DeleteRestaurantTable elimina un registro de restaurant_table por su ID
func (h *RestaurantTableHandler) DeleteRestaurantTable(c echo.Context) error {
	// Obtener el ID del parámetro de la ruta
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid restaurant_table ID"})
	}

	// Llamar al método de la aplicación para eliminar el registro
	err = h.app.DeleteRestaurantTable(c.Request().Context(), int32(id))
	if err != nil {
		// Si no se encuentra el registro o hay otro error, manejar el error apropiadamente
		if err == sql.ErrNoRows {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "RestaurantTable not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Responder con un mensaje de éxito
	return c.JSON(http.StatusOK, map[string]string{"message": "Registro eliminado con éxito"})
}

// UpdateRestaurantTable maneja la actualización de registros en restaurant_table
func (h *RestaurantTableHandler) UpdateRestaurantTable(c echo.Context) error {
	// Obtener el ID del parámetro de la URL
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid restaurant_table ID"})
	}

	var req dto.UpdateRestaurantTableRequest
	// Decodificar el cuerpo de la solicitud en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	// Llamar al método de la aplicación para actualizar el registro
	err = h.app.UpdateRestaurantTable(c.Request().Context(), int32(id), req.Number, req.StatusID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Registro actualizado con éxito"})
}
