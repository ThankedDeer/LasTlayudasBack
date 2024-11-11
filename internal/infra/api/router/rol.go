package router

import (
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
)

type IRoleRouter interface {
	RoleResource(g *echo.Group)
}

type RoleRouter struct {
	handler handler.IRoleHandler
}

func NewRoleRouter(handler handler.IRoleHandler) IRoleRouter {
	return &RoleRouter{
		handler: handler,
	}
}

func (u *RoleRouter) RoleResource(g *echo.Group) {
	groupPath := g.Group("/role")
	groupPath.POST("/", u.handler.CreateRole)
	groupPath.GET("/", u.handler.GetAllRoles)
	groupPath.GET("/:id", u.handler.GetRoleByID)
	groupPath.PUT("/:id", u.handler.UpdateRole)
	groupPath.DELETE("/:id", u.handler.DeleteRole)
}
