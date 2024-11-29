package router

import (
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
)

type IPermissionRouter interface {
	PermissionResource(g *echo.Group)
}

type PermissionRouter struct {
	handler handler.IPermissionHandler
}

func NewPermissionRouter(handler handler.IPermissionHandler) IPermissionRouter {
	return &PermissionRouter{
		handler: handler,
	}
}

func (r *PermissionRouter) PermissionResource(g *echo.Group) {
	groupPath := g.Group("/permissions")
	groupPath.POST("/", r.handler.CreatePermission)
	groupPath.GET("/", r.handler.GetAllPermissions)
	groupPath.PUT("/:id", r.handler.UpdatePermission)
	groupPath.DELETE("/:id", r.handler.DeletePermission)
	groupPath.GET("/:id", r.handler.GetPermissionByID)
}
