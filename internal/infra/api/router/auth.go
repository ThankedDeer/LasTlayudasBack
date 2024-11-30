package router

import (
	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/infra/api/handler"
)

type IAuthRouter interface {
	AuthResource(g *echo.Group)
}

type AuthRouter struct {
	handler handler.IAuthHandler
}

func NewAuthRouter(handler handler.IAuthHandler) IAuthRouter {
	return &AuthRouter{
		handler: handler,
	}
}

func (u *AuthRouter) AuthResource(g *echo.Group) {
	groupPath := g.Group("/auth")
	groupPath.POST("/login", u.handler.Login)
}
