package api

import (
	"fmt"
	"github/thankeddeer/lastlayudas/config"
	"github/thankeddeer/lastlayudas/internal/infra/api/router"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	cfg    config.Config
	engine *echo.Echo
	//userRouter router.IUserRouter
	ProductRouter  router.IProductRouter
	CategoryRouter router.ICategoryRouter
	ProviderRouter router.IProviderRouter
	RoleRouter     router.IRoleRouter
	PermissionRouter router.IPermissionRouter
}

func NewServer(
	cfg config.Config,
	engine *echo.Echo,
	ProductRouter router.IProductRouter,
	CategoryRouter router.ICategoryRouter,
	ProviderRouter router.IProviderRouter,
	RoleRouter router.IRoleRouter,
	//userRouter router.IUserRouter,
	PermissionRouter router.IPermissionRouter,
) *Server {
	return &Server{
		cfg:            cfg,
		engine:         engine,
		ProductRouter:  ProductRouter,
		CategoryRouter: CategoryRouter,
		ProviderRouter: ProviderRouter,
		RoleRouter:     RoleRouter,
		//userRouter: userRouter,
		PermissionRouter: PermissionRouter,
	}

}

func (s *Server) BuildServer() {
	s.engine.Use(middleware.CORS())
	s.engine.Use(middleware.Logger())
	s.engine.Use(middleware.Recover())

	basePath := s.engine.Group("/api")
	s.ProductRouter.ProductResource(basePath)
	s.CategoryRouter.CategoryResource(basePath)
	s.ProviderRouter.ProviderResource(basePath)
	//s.userRouter.UserResource(basePath)
	s.PermissionRouter.PermissionResource(basePath)
	fmt.Println(basePath)

}

func (s *Server) Run() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error during server initialization:", err)
		}
	}()

	if err := s.engine.Start(fmt.Sprintf(":%s", s.cfg.ServerAddress)); err != nil {
		log.Printf("Error starting server: %v", err)
		return fmt.Errorf("error starting server: %w", err)
	}

	return nil
}
