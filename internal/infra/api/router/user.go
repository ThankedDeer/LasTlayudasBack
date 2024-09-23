package router

import (
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
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
	groupPath := g.Group("/users")
	groupPath.POST("/", u.handler.CreateUser)
	groupPath.POST("/Testimonials", u.handler.CreateTestimonial)
	groupPath.GET("/Testimonials", u.handler.GetTestimonials)
	groupPath.GET("/", u.handler.GetUsers)
	groupPath.PUT("/", u.handler.UpdateUser)
}
