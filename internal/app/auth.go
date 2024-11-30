package app

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

// AuthApp contiene la lógica de negocio para la autenticación.
type AuthApp struct {
	store       *sqlc.Store
	jwtSecret   string
	tokenExpiry time.Duration
}

// NewAuthApp crea una nueva instancia de AuthApp.
func NewAuthApp(store *sqlc.Store, jwtSecret string, tokenExpiry time.Duration) *AuthApp {
	return &AuthApp{
		store:       store,
		jwtSecret:   jwtSecret,
		tokenExpiry: tokenExpiry,
	}
}

// Login realiza la validación de credenciales y devuelve un token JWT.
func (a *AuthApp) Login(ctx context.Context, email, password string) (string, error) {
	user, err := a.store.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("invalid credentials")
		}
		return "", err
	}

	if !user.IsActive.Bool {
		return "", errors.New("account is inactive")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := a.generateToken(user.UserID, user.RoleID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateToken genera un JWT para un usuario autenticado.
func (a *AuthApp) generateToken(userID int32, roleID int32) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role_id": roleID,
		"exp":     time.Now().Add(a.tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.jwtSecret))
}

// HashPassword genera un hash de contraseña usando bcrypt.
func (a *AuthApp) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
