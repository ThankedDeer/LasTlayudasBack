package router

import (
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
)

type IProviderRouter interface {
	ProviderResource(g *echo.Group)
}

type ProviderRouter struct {
	handler handler.IProviderHandler
}

func NewProviderRouter(handler handler.IProviderHandler) IProviderRouter {
	return &ProviderRouter{
		handler: handler,
	}
}

func (u *ProviderRouter) ProviderResource(g *echo.Group) {
	groupPath := g.Group("/provider")
	groupPath.POST("/", u.handler.CreateProvider)
	groupPath.GET("/", u.handler.GetAllProvider)
	groupPath.PUT("/", u.handler.UpdateProvider)
	groupPath.DELETE("/:id", u.handler.DeleteProviderHandler)
}
