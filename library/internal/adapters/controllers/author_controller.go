package controllers

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
	"github.com/gin-gonic/gin"
)

// Primary Adapter
type AuthorController struct {
	authorService ports.AuthorService
	engine        *gin.Engine
}

func NewAuthorController(authorService ports.AuthorService, engine *gin.Engine) AuthorController {
	return AuthorController{
		authorService: authorService,
		engine:        engine,
	}
}

func (ac AuthorController) InitAuthorRouters() {
	group := ac.engine.Group("/api/v1/author")

	group.GET("/find/:id", ac.authorService.FindAuthor)
}
