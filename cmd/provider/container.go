package provider

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"

	"github/thankeddeer/lastlayudas/config"
	"github/thankeddeer/lastlayudas/internal/app"
	"github/thankeddeer/lastlayudas/internal/infra/api"
	"github/thankeddeer/lastlayudas/internal/infra/api/handler"
	"github/thankeddeer/lastlayudas/internal/infra/api/router"
	"github/thankeddeer/lastlayudas/internal/store/sqlc"
)

type Container struct{}

func NewProvider() *Container {
	return &Container{}
}

func (c *Container) Build() *api.Server {

	config, err := config.LoadConfig("./")

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

	// Parámetros para la autenticación
	jwtSecret := "my_secret_key"  // Debería ser un secreto fuerte y seguro
	tokenExpiry := 24 * time.Hour // Ejemplo de duración del token

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

	UserService := app.NewUserApp(store)
	UserHandler := handler.NewUserHandler(UserService)
	UserRouter := router.NewUserRouter(UserHandler)

	AuthService := app.NewAuthApp(store, jwtSecret, tokenExpiry)
	AuthHandler := handler.NewAuthHandler(AuthService)
	AuthRouter := router.NewAuthRouter(AuthHandler)

	server := api.NewServer(
		config,
		engine,
		ProductRouter,
		CategoryRouter,
		ProviderRouter,
		RoleRouter,
		UserRouter,
		AuthRouter,
	)
	server.BuildServer()
	return server

}
