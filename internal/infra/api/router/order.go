package router

import (
	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/infra/api/handler"
)

type IOrderRouter interface {
	OrderResource(g *echo.Group)
}

type OrderRouter struct {
	handler handler.IOrderHandler
}

func NewOrderRouter(handler handler.IOrderHandler) IOrderRouter {
	return &OrderRouter{
		handler: handler,
	}
}

func (u *OrderRouter) OrderResource(g *echo.Group) {
	groupPath := g.Group("/order")
	groupPath.POST("/", u.handler.CreateOrder)      // Crear una nueva orden
	groupPath.GET("/", u.handler.GetAllOrders)      // Obtener todas las órdenes
	groupPath.PUT("/:id", u.handler.UpdateOrder)    // Actualizar una orden por su ID
	groupPath.DELETE("/:id", u.handler.DeleteOrder) // Eliminar una orden por su ID
	groupPath.GET("/:id", u.handler.GetOrderByID)   // Obtener una orden por su ID
}
