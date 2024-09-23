package handler

import (
	"encoding/json"
	"errors"
	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/domain/dto"
	"github/thankeddeer/lastlayudas/store/sqlc"
	"net/http"
	"unicode"

	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
)

type IUserHandler interface {
	CreateUser(c echo.Context) error
	GetUsers(c echo.Context) error
	UpdateUser(c echo.Context) error
	CreateTestimonial(c echo.Context) error
	GetTestimonials(c echo.Context) error
}

type UserHandler struct {
	app app.UserApp
}

func NewUserHandler(app app.UserApp) IUserHandler {
	return &UserHandler{
		app: app,
	}

}

func (u *UserHandler) CreateUser(c echo.Context) error {
	var req dto.CreateUserRequest

	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := validatePassword(req.Password); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := u.app.CreateUser(req)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				return c.JSON(http.StatusForbidden, err)
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)

}
func (u *UserHandler) GetUsers(c echo.Context) error {

	users, err := u.app.GetUsers()

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				return c.JSON(http.StatusForbidden, err)
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)

}

func (u *UserHandler) UpdateUser(c echo.Context) error {
	var req sqlc.UpdatUserParams

	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := u.app.UpdateUser(req)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				return c.JSON(http.StatusForbidden, err)
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)

}

func (u *UserHandler) CreateTestimonial(c echo.Context) error {
	var req dto.CreateTestimonial

	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user, err := u.app.CreateTestimonial(req)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				return c.JSON(http.StatusForbidden, err)
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)

}

func (u *UserHandler) GetTestimonials(c echo.Context) error {

	users, err := u.app.GetTestimonial()

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			// log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation", "foreign_key_violation":
				return c.JSON(http.StatusForbidden, err)
			}
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)

}
func validatePassword(pass string) error {
	var (
		hasMinLen    = false
		hasUppercase = false
		hasLowercase = false
		hasSpecial   = false
		hasNoSeqNum  = true
		hasNoSeqLet  = true
	)

	if len(pass) >= 8 {
		hasMinLen = true
	}

	for i, char := range pass {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}

		// Verificar números consecutivos
		if i > 0 && unicode.IsDigit(char) && unicode.IsDigit(rune(pass[i-1])) {
			if char == rune(pass[i-1])+1 {
				hasNoSeqNum = false
			}
		}

		// Verificar letras consecutivas
		if i > 0 && unicode.IsLetter(char) && unicode.IsLetter(rune(pass[i-1])) {
			if unicode.ToLower(char) == unicode.ToLower(rune(pass[i-1]))+1 {
				hasNoSeqLet = false
			}
		}
	}

	if !hasMinLen {
		return errors.New("la contraseña debe tener al menos 8 caracteres")
	}
	if !hasUppercase {
		return errors.New("la contraseña debe tener al menos una mayúscula")
	}
	if !hasLowercase {
		return errors.New("la contraseña debe tener al menos una minúscula")
	}
	if !hasSpecial {
		return errors.New("la contraseña debe tener al menos un carácter especial")
	}
	if !hasNoSeqNum {
		return errors.New("la contraseña no debe contener números consecutivos")
	}
	if !hasNoSeqLet {
		return errors.New("la contraseña no debe contener letras consecutivas")
	}

	return nil
}
