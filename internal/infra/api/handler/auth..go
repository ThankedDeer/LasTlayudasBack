package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
)

// IAuthHandler define los métodos para el handler de autenticación
type IAuthHandler interface {
	Login(ctx echo.Context) error
}

// AuthHandler implementa la interfaz IAuthHandler
type AuthHandler struct {
	app *app.AuthApp
}

// NewAuthHandler crea una nueva instancia de AuthHandler
func NewAuthHandler(authService *app.AuthApp) IAuthHandler {
	return &AuthHandler{
		app: authService,
	}
}

// Login maneja la autenticación de usuarios
func (h *AuthHandler) Login(ctx echo.Context) error {
	var req dto.LoginRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	// Llamar a la lógica de negocio de AuthApp
	token, err := h.app.Login(ctx.Request().Context(), req.Email, req.Password)
	if err != nil {
		if err.Error() == "invalid credentials" || err.Error() == "account is inactive" {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": err.Error()})
		}
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": "Server error"})
	}

	// Respuesta exitosa con el token generado
	return ctx.JSON(http.StatusOK, dto.LoginResponse{
		Message: "Login successful",
		Token:   token,
	})
}
