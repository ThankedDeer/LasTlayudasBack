package handler

import (
	"database/sql"
	"encoding/json"
	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IPermissionHandler interface {
	CreatePermission(c echo.Context) error
	GetAllPermissions(c echo.Context) error
	UpdatePermission(c echo.Context) error
	DeletePermission(c echo.Context) error
	GetPermissionByID(c echo.Context) error
}

type PermissionHandler struct {
	app app.PermissionApp
}

func NewPermissionHandler(app app.PermissionApp) IPermissionHandler {
	return &PermissionHandler{
		app: app,
	}
}

// CreatePermission maneja la creación de un nuevo permiso
func (p *PermissionHandler) CreatePermission(c echo.Context) error {
	var req dto.CreatePermissionRequest

	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos: " + err.Error()})
	}

	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "El campo 'name' es obligatorio."})
	}

	permission, err := p.app.CreatePermission(c.Request().Context(), req.Name, req.Description)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, permission)
}

// GetAllPermissions obtiene todos los permisos
func (p *PermissionHandler) GetAllPermissions(c echo.Context) error {
	permissions, err := p.app.GetAllPermissions(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Permisos no encontrados"})
	}

	return c.JSON(http.StatusOK, permissions)
}

// UpdatePermission maneja la actualización de un permiso
func (p *PermissionHandler) UpdatePermission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de permiso inválido"})
	}

	var req dto.UpdatePermissionRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	description := sql.NullString{Valid: req.Description != nil}
	if req.Description != nil {
		description.String = *req.Description
	}

	updateParams := sqlc.UpdatePermissionParams{
		PermissionID: int32(id),
		Name:         req.Name,
		Description:  description,
	}

	if err := p.app.UpdatePermission(c.Request().Context(), updateParams); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Permiso actualizado con éxito"})
}

// DeletePermission maneja la eliminación de un permiso
func (p *PermissionHandler) DeletePermission(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de permiso inválido"})
	}

	if err := p.app.DeletePermission(c.Request().Context(), int32(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Permiso eliminado correctamente"})
}

// GetPermissionByID obtiene un permiso por su ID
func (p *PermissionHandler) GetPermissionByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de permiso inválido"})
	}

	permission, err := p.app.GetPermissionByID(c.Request().Context(), int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Permiso no encontrado"})
	}

	return c.JSON(http.StatusOK, permission)
}
