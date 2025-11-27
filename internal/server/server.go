package server

import (
	"fmt"
	"test-1/internal/configs"
	"test-1/internal/database"
	"test-1/internal/handlers"
	"test-1/internal/repositories"
	"test-1/internal/routers"
	"test-1/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Server struct {
	l      zerolog.Logger
	config *configs.Config
}

func CreateServer(logger zerolog.Logger, config *configs.Config) *Server {
	return &Server{
		l:      logger,
		config: config,
	}
}

func (s *Server) Run() {
	address := s.config.Server.Address
	port := s.config.Server.Port

	// Start Database
	database.CreateDatabaseConnection(s.config)

	// Init Repositories, Services, Handlers here if needed
	userRepository := repositories.NewUserRepository(database.Db)

	userService := services.Constructor(userRepository)

	userHandler := handlers.Constructor(userService)

	// Build router
	router := gin.Default()

	// Register public endpoint
	routers.RegisterPublicEndpoint(router)
	routers.RegisterUserPublicEndpoint(router, userHandler)

	s.l.Info().Msgf("Starting server at %s:%d", address, port)
	s.l.Info().Msgf("Starting server at http://127.0.0.1:%d", port)

	if err := router.Run(fmt.Sprintf("%s:%d", address, port)); err != nil {
		s.l.Error().Err(err).Msg("Failed to start server")
	}
}
