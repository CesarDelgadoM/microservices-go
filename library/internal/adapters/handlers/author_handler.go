package handlers

import (
	"net/http"

	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
	"github.com/gin-gonic/gin"
)

// Primary Adapter
type AuthorHandler struct {
	authorService ports.AuthorService
	engine        *gin.Engine
}

func NewAuthorController(authorService ports.AuthorService, engine *gin.Engine) AuthorHandler {
	return AuthorHandler{
		authorService: authorService,
		engine:        engine,
	}
}

func (ah AuthorHandler) InitAuthorRouters() {
	group := ah.engine.Group("/api/v1/author")

	group.GET("/find/:id", ah.FindAuthor)
	group.POST("/create", ah.CreateAuthor)
}

func (ah AuthorHandler) FindAuthor(ctx *gin.Context) {
	ah.authorService.FindAuthor(0)
}

func (ah AuthorHandler) CreateAuthor(ctx *gin.Context) {
	var author entities.Author

	if err := ctx.BindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest,
			domain.ResponseError(err.Error(), http.StatusBadRequest))
		return
	}

	err := ah.authorService.CreateAuthor(&author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			domain.ResponseError(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, domain.ResponseData(author, "Author created",
		http.StatusOK))
}
