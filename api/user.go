package api

import (
	db "github/thankeddeer/lastlayudas/db/sqlc"
	"net/http"

	"github.com/labstack/echo/v4"
)

type createUserRequest struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

func (server *Server) createUser(ctx echo.Context) error {

	var req createUserRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, errorResponse(err))

	}

	arg := db.CreateUserParams{
		Email:     req.Email,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Password:  req.Password,
	}

	user, err := server.store.CreateUser(ctx.Request().Context(), arg)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, errorResponse(err))

	}

	return ctx.JSON(http.StatusOK, user)

}
