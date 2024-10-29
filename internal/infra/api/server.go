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
	cfg        config.Config
	engine     *echo.Echo
	userRouter router.IUserRouter
}

func NewServer(
	cfg config.Config,
	engine *echo.Echo,
	userRouter router.IUserRouter,
) *Server {
	return &Server{
		cfg:        cfg,
		engine:     engine,
		userRouter: userRouter,
	}

}

func (s *Server) BuildServer() {
	s.engine.Use(middleware.CORS())
	s.engine.Use(middleware.Logger())
	s.engine.Use(middleware.Recover())

	basePath := s.engine.Group("/api")
	s.userRouter.UserResource(basePath)
	
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
