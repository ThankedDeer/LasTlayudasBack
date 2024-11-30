package router

import (
	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/infra/api/handler"
)

type IUserRouter interface {
	UserResource(g *echo.Group)
}

type UserRouter struct {
	handler handler.IUserHandler
}

func NewUserRouter(handler handler.IUserHandler) IUserRouter {
	return &UserRouter{
		handler: handler,
	}
}

func (u *UserRouter) UserResource(g *echo.Group) {
	groupPath := g.Group("/user")
	groupPath.POST("/", u.handler.CreateUser)
	groupPath.GET("/", u.handler.GetAllUsers)
	groupPath.PUT("/:id", u.handler.UpdateUser)
	groupPath.DELETE("/:id", u.handler.DeleteUser)
}
