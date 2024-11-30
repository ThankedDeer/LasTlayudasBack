package router

import (
	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/internal/infra/api/handler"
)

// IRestaurantTableRouter define las operaciones para manejar los recursos de restaurant_table
type IRestaurantTableRouter interface {
	RestaurantTableResource(g *echo.Group)
}

// RestaurantTableRouter implementa la interfaz IRestaurantTableRouter
type RestaurantTableRouter struct {
	handler handler.IRestaurantTableHandler
}

// NewRestaurantTableRouter crea una nueva instancia de RestaurantTableRouter
func NewRestaurantTableRouter(handler handler.IRestaurantTableHandler) IRestaurantTableRouter {
	return &RestaurantTableRouter{
		handler: handler,
	}
}

// RestaurantTableResource define las rutas para los recursos de restaurant_table
func (u *RestaurantTableRouter) RestaurantTableResource(g *echo.Group) {
	groupPath := g.Group("/restaurant_table")
	groupPath.POST("/", u.handler.CreateRestaurantTable)      // Ruta para crear una nueva mesa
	groupPath.GET("/", u.handler.GetAllRestaurantTables)      // Ruta para obtener todas las mesas
	groupPath.PUT("/:id", u.handler.UpdateRestaurantTable)    // Ruta para actualizar una mesa específica
	groupPath.DELETE("/:id", u.handler.DeleteRestaurantTable) // Ruta para eliminar una mesa específica
}
