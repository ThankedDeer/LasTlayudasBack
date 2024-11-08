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

type IRoleHandler interface {
	CreateRole(c echo.Context) error
	GetAllRoles(c echo.Context) error
	UpdateRole(c echo.Context) error
	DeleteRole(c echo.Context) error
	GetRoleByID(c echo.Context) error
}

type RoleHandler struct {
	app app.RoleApp
}

func NewRoleHandler(app app.RoleApp) IRoleHandler {
	return &RoleHandler{
		app: app,
	}
}

func (u *RoleHandler) CreateRole(c echo.Context) error {
	var req dto.CreateRoleRequest

	// Decodificar el cuerpo JSON en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos: " + err.Error()})
	}

	// Validación: asegurarse de que 'Name' no esté vacío
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'nombre' es obligatorio."})
	}

	role, err := u.app.CreateRole(c.Request().Context(), req.Name, req.Description)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, role)
}

func (u *RoleHandler) GetAllRole(c echo.Context) error {
	roles, err := u.app.GetAllRoles(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Roles no encontrados"})
	}

	return c.JSON(http.StatusOK, roles)
}

func (u *RoleHandler) UpdateRole(c echo.Context) error {
	// Convertir el ID de categoría de la URL a entero
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de rol inválido"})
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

	// Crear los parámetros de actualización usando los datos decodificados
	updateParams := sqlc.UpdateRoleParams{
		RoleID:      int32(id),
		Name:        req.Name,
		Description: description,
	}

	// Ejecutar la actualización de la categoría en la base de datos
	if err := u.app.UpdateRole(c.Request().Context(), updateParams); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Respuesta de éxito si la actualización fue exitosa
	return c.JSON(http.StatusOK, map[string]string{"message": "Rol actualizado con éxito"})
}

func (u *RoleHandler) DeleteRole(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de rol inválido"})
	}

	if err := u.app.DeleteRole(c.Request().Context(), int32(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Responder con un mensaje de éxito y código 200 OK
	return c.JSON(http.StatusOK, map[string]string{"message": "Categoría eliminada correctamente"})
}

func (u *RoleHandler) GetRoleByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de rol inválido"})
	}

	role, err := u.app.GetRoleByID(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Rol no encontrado"})
	}

	return c.JSON(http.StatusOK, role)
}
