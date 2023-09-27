package server

import (
	"github.com/CesarDelgadoM/microservices-go/library/database"
	"github.com/CesarDelgadoM/microservices-go/library/internal/adapters/handlers"
	"github.com/CesarDelgadoM/microservices-go/library/internal/adapters/repositories"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/services"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db     database.IDatabase
	engine *gin.Engine
}

func NewServer(db database.IDatabase, engine *gin.Engine) *Server {
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

	// init handlers
	authorHandler := handlers.NewAuthorController(authorService, server.engine)

	// init routes
	authorHandler.InitAuthorRouters()

	return server.engine.Run(addr)
}
