package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
)

type IUserHandler interface {
	CreateUser(c echo.Context) error
	GetAllUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	DeleteUser(c echo.Context) error
	GetUserByID(c echo.Context) error
}

type UserHandler struct {
	app app.UserApp
}

func NewUserHandler(app app.UserApp) IUserHandler {
	return &UserHandler{
		app: app,
	}
}

// CreateUser maneja la creación de un nuevo usuario
func (u *UserHandler) CreateUser(c echo.Context) error {
	var req dto.CreateUserRequest

	// Decodificar el cuerpo JSON en la estructura `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos: " + err.Error()})
	}

	// Validación: asegurarse de que los campos requeridos no estén vacíos
	if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Todos los campos son obligatorios."})
	}

	// Crear el usuario a través de la lógica de aplicación
	user, err := u.app.CreateUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, user)
}

// GetAllUsers obtiene todos los usuarios
func (u *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := u.app.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Usuarios no encontrados"})
	}

	return c.JSON(http.StatusOK, users)
}

// UpdateUser maneja la actualización de un usuario
func (u *UserHandler) UpdateUser(c echo.Context) error {
	// Convertir el ID del usuario de la URL a entero
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de usuario inválido"})
	}

	var req dto.UpdateUserRequest
	// Decodificar el JSON del cuerpo de la solicitud en `req`
	if err := json.NewDecoder(c.Request().Body).Decode(&req); err != nil {
		log.Println("Error de decodificación:", err) // Log para el error de decodificación
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	// Actualizar el usuario usando la lógica de aplicación
	req.UserID = int32(id)
	if err := u.app.UpdateUser(req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Respuesta de éxito si la actualización fue exitosa
	return c.JSON(http.StatusOK, map[string]string{"message": "Usuario actualizado con éxito"})
}

// DeleteUser maneja la eliminación de un usuario por su ID
func (u *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de usuario inválido"})
	}

	if err := u.app.DeleteUser(int32(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Responder con un mensaje de éxito y código 200 OK
	return c.JSON(http.StatusOK, map[string]string{"message": "Usuario eliminado correctamente"})
}

// GetUserByID obtiene un usuario por su ID
func (u *UserHandler) GetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID de usuario inválido"})
	}

	user, err := u.app.GetUserByID(int32(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Usuario no encontrado"})
	}

	return c.JSON(http.StatusOK, user)
}
