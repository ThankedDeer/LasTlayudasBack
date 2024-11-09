package provider

import (
	"database/sql"
	"fmt"
	"github/thankeddeer/lastlayudas/config"
	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/infra/api"
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"
	"github/thankeddeer/lastlayudas/internal/infra/api/router"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
	"log"

	"github.com/labstack/echo/v4"
)

type Container struct{}

func NewProvider() *Container {
	return &Container{}
}

func (c *Container) Build() *api.Server {

	config, err := config.LoadConfig("../../.")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	fmt.Println(conn)
	engine := echo.New()

	store := sqlc.NewStore(conn)

	ProductService := app.NewProductApp(store)
	ProductHandler := handler.NewProductHandler(ProductService)
	ProductRouter := router.NewProductRouter(ProductHandler)

	CategoryService := app.NewCategoryApp(store)
	CategoryHandler := handler.NewCategoryHandler(CategoryService)
	CategoryRouter := router.NewCategoryRouter(CategoryHandler)

	ProviderService := app.NewProviderApp(store)
	ProviderHandler := handler.NewProviderHandler(ProviderService)
	ProviderRouter := router.NewProviderRouter(ProviderHandler)
	
	RoleService := app.NewRoleApp(store)
	RoleHandler := handler.NewRoleHandler(RoleService)
	RoleRouter := router.NewRoleRouter(RoleHandler)

	server := api.NewServer(
		config,
		engine,
		ProductRouter,
		CategoryRouter,
		ProviderRouter,
		RoleRouter
	)
	server.BuildServer()
	return server

}
