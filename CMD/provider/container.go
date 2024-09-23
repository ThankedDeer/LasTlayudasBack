package provider

import (
	"database/sql"
	"github/thankeddeer/lastlayudas/config"
	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/infra/api"
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"
	"github/thankeddeer/lastlayudas/internal/infra/api/router"
	"github/thankeddeer/lastlayudas/store/sqlc"
	"log"

	"github.com/labstack/echo/v4"

)

type Container struct{}

func NewProvider() *Container {
	return &Container{}
}

func (c *Container) Build() *api.Server {

	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	engine := echo.New()

	//son todas los repositorios que genera sqlc
	store := sqlc.NewStore(conn)

	userService := app.NewUserApp(store)
	userHandler := handler.NewUserHandler(userService)
	useRouter := router.NewUserRouter(userHandler)

	server := api.NewServer(
		config,
		engine,
		useRouter,
	)
	server.BuildServer()
	return server

}
