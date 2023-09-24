package server

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/adapters/controllers"
	"github.com/CesarDelgadoM/microservices-go/library/internal/adapters/repositories"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/services"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     ports.Database
	engine *gin.Engine
}

func NewServer(db ports.Database, engine *gin.Engine) *Server {
	return &Server{
		db:     db,
		engine: engine,
	}
}

func (server Server) Start(addr string) error {
	// init repositories
	authorRepository := repositories.NewAuthorRepository(server.db)

	// init services
	authorService := services.NewAuthorService(authorRepository)

	// init controllers
	controllers := controllers.NewAuthorController(authorService, server.engine)

	// init routes
	controllers.InitAuthorRouters()

	return server.engine.Run(addr)
}
