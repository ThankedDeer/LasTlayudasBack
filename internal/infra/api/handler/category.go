package handler

import (
	"database/sql"
	"encoding/json"
	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICategoryHandler interface {
	CreateCategory(c echo.Context) error
	GetAllCategories(c echo.Context) error
	UpdateCategory(c echo.Context) error
	DeleteCategory(c echo.Context) error
	GetCategoryByID(c echo.Context) error
}

type CategoryHandler struct {
	app app.CategoryApp
}

func NewCategoryHandler(app app.CategoryApp) ICategoryHandler {
	return &CategoryHandler{
		app: app,
	}
}

// CreateCategory maneja la creación de una nueva categoría
// @Summary Create a new category
// @Description Create a new category with the input payload
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryRequest true "Category to create"
// @Success 201 {object} sqlc.Category
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/category/ [post]
func (u *CategoryHandler) CreateCategory(c echo.Context) error {
	var req dto.CreateCategoryRequest

	// Decodificar el cuerpo JSON en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos: " + err.Error()})
	}

	// Validación: asegurarse de que 'Name' no esté vacío
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'name' es obligatorio."})
	}

	// Asignar true a 'is_active' si no se envía en la solicitud
	if c.Request().Header.Get("Content-Type") == "application/json" && !req.Is_active {
		req.Is_active = true
	}

	category, err := u.app.CreateCategory(c.Request().Context(), req.Name, req.Description, req.Is_active)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, category)
}

// GetAllCategories obtiene todas las categorías
// @Summary Get all categories
// @Description Retrieve all categories
// @Tags categories
// @Produce json
// @Success 200 {array} sqlc.Category
// @Failure 404 {object} map[string]string
// @Router /api/category/ [get]
func (u *CategoryHandler) GetAllCategories(c echo.Context) error {
	categories, err := u.app.GetAllCategories(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Categorías no encontradas"})
	}

	return c.JSON(http.StatusOK, categories)
}

// UpdateCategory maneja la actualización de una categoría
// @Summary Update an existing category
// @Description Update an existing category with the input payload
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body dto.UpdateCategoryRequest true "Category to update"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/category/{id} [put]
func (u *CategoryHandler) UpdateCategory(c echo.Context) error {
	// Convertir el ID de categoría de la URL a entero
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de categoría inválido"})
	}

	var req dto.UpdateCategoryRequest
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

	// Convertir *bool a sql.NullBool para el campo is_active
	is_active := sql.NullBool{Bool: req.Is_active, Valid: true}

	// Crear los parámetros de actualización usando los datos decodificados
	updateParams := sqlc.UpdateCategoryParams{
		CategoryID:  int32(id),
		Name:        req.Name,
		Description: description,
		IsActive:    is_active,
	}

	// Ejecutar la actualización de la categoría en la base de datos
	if err := u.app.UpdateCategory(c.Request().Context(), updateParams); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Respuesta de éxito si la actualización fue exitosa
	return c.JSON(http.StatusOK, map[string]string{"message": "Categoría actualizada con éxito"})
}

// DeleteCategory maneja la eliminación de una categoría por su ID
// @Summary Delete a category
// @Description Delete a category by its ID
// @Tags categories
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/category/{id} [delete]
func (u *CategoryHandler) DeleteCategory(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de categoría inválido"})
	}

	if err := u.app.DeleteCategory(c.Request().Context(), int32(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Responder con un mensaje de éxito y código 200 OK
	return c.JSON(http.StatusOK, map[string]string{"message": "Categoría eliminada correctamente"})
}

// GetCategoryByID obtiene una categoría por su ID
// @Summary Get a category by ID
// @Description Retrieve a category by its ID
// @Tags categories
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} sqlc.Category
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/category/{id} [get]
func (u *CategoryHandler) GetCategoryByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de categoría inválido"})
	}

	category, err := u.app.GetCategoryByID(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Categoría no encontrada"})
	}

	return c.JSON(http.StatusOK, category)
}
