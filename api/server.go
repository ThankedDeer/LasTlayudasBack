package api

import (
	db "github/thankeddeer/lastlayudas/db/sqlc"

	"github.com/labstack/echo/v4"
)

type Server struct {
	store  *db.Store
	router *echo.Echo
}

func NewServer(store *db.Store) *Server {

	server := &Server{store: store}
	router := echo.New()

	router.POST("/users", server.createUser)

	server.router = router

	return server

}
func (server *Server) Start(address string) error {
	return server.router.Start(address)
}

func errorResponse(err error) echo.HTTPError {

	return echo.HTTPError{
		Message: err.Error(),
	}
}
