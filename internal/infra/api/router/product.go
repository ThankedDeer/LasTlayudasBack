package router

import (
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
)

type IProductRouter interface {
	ProductResource(g *echo.Group)
}

type ProductRouter struct {
	handler handler.IProductHandler
}

func NewProductRouter(handler handler.IProductHandler) IProductRouter {
	return &ProductRouter{
		handler: handler,
	}
}

func (u *ProductRouter) ProductResource(g *echo.Group) {
	groupPath := g.Group("/product")
	groupPath.POST("/", u.handler.CreateProduct)
	groupPath.GET("/", u.handler.GetAllProducts)
	groupPath.PUT("/", u.handler.UpdateProduct)
	groupPath.DELETE("/", u.handler.UpdateProduct)
}
