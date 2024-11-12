package router

import (
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"

	"github.com/labstack/echo/v4"
)

type IOrderStatusRouter interface {
	OrderStatusResource(g *echo.Group)
}

type OrderStatusRouter struct {
	handler handler.IOrderStatusHandler
}

func NewOrdeStatusRouter(handler handler.IOrderStatusHandler) IOrderStatusRouter {
	return &OrderStatusRouter{
		handler: handler,
	}
}

func (u *OrderStatusRouter) OrderStatusResource(g *echo.Group) {
	groupPath := g.Group("/order-status")
	groupPath.POST("/", u.handler.CreateOrderStatus)
	groupPath.GET("/", u.handler.GetAllOrderStatus)
	groupPath.GET("/:id", u.handler.GetOrderStatusByID)
	groupPath.PUT("/:id", u.handler.UpdateOrderStatus)
	groupPath.DELETE("/:id", u.handler.DeleteOrderStatus)
}
