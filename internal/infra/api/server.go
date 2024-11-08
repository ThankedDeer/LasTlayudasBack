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
	RoleRouter     router.IRoleRouter
}

func NewServer(
	cfg config.Config,
	engine *echo.Echo,
	ProductRouter router.IProductRouter,
	CategoryRouter router.ICategoryRouter,
	RoleRouter router.IRoleRouter,
	//userRouter router.IUserRouter,
) *Server {
	return &Server{
		cfg:            cfg,
		engine:         engine,
		ProductRouter:  ProductRouter,
		CategoryRouter: CategoryRouter,
		RoleRouter:     RoleRouter,
		//userRouter: userRouter,
	}

}

func (s *Server) BuildServer() {
	s.engine.Use(middleware.CORS())
	s.engine.Use(middleware.Logger())
	s.engine.Use(middleware.Recover())

	basePath := s.engine.Group("/api")
	s.ProductRouter.ProductResource(basePath)
	s.CategoryRouter.CategoryResource(basePath)
	s.RoleRouter.RoleResource(basePath)
	//s.userRouter.UserResource(basePath)
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
