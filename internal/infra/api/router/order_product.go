package router

import (
	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/infra/api/handler"
)

type IOrderProductRouter interface {
	OrderProductResource(g *echo.Group)
}

type OrderProductRouter struct {
	handler handler.IOrderProductHandler
}

func NewOrderProductRouter(handler handler.IOrderProductHandler) IOrderProductRouter {
	return &OrderProductRouter{
		handler: handler,
	}
}

func (u *OrderProductRouter) OrderProductResource(g *echo.Group) {
	groupPath := g.Group("/order_product")
	groupPath.POST("/", u.handler.CreateOrderProduct)
	groupPath.GET("/:id", u.handler.GetOrderProductByID)
	groupPath.PUT("/:id", u.handler.UpdateOrderProduct)
	groupPath.DELETE("/:id", u.handler.DeleteOrderProduct)
}
