package provider

import (
	"database/sql"
	"fmt"
	"github/thankeddeer/lastlayudas/config"
	"github/thankeddeer/lastlayudas/internal/infra/api"
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
	fmt.Println(conn)
	engine := echo.New()

	server := api.NewServer(
		config,
		engine,
	)
	server.BuildServer()
	return server

}
