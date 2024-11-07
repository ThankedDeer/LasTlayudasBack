package router

import (
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
)

type ICategoryRouter interface {
	CategoryResource(g *echo.Group)
}

type CategoryRouter struct {
	handler handler.ICategoryHandler
}

func NewCategoryRouter(handler handler.ICategoryHandler) ICategoryRouter {
	return &CategoryRouter{
		handler: handler,
	}
}

func (u *CategoryRouter) CategoryResource(g *echo.Group) {
	groupPath := g.Group("/category")
	groupPath.POST("/", u.handler.CreateCategory)
	groupPath.GET("/", u.handler.GetAllCategories)
	groupPath.PUT("/:id", u.handler.UpdateCategory)
	groupPath.DELETE("/:id", u.handler.DeleteCategory)
}
