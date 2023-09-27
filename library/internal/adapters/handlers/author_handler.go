package handlers

import (
	"fmt"
	"net/http"

	"github.com/CesarDelgadoM/microservices-go/library/internal/core/common"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
	"github.com/gin-gonic/gin"
)

// Primary Adapter
type AuthorHandler struct {
	authorService ports.IAuthorService
	engine        *gin.Engine
}

func NewAuthorController(authorService ports.IAuthorService, engine *gin.Engine) AuthorHandler {
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

	if err := ctx.ShouldBindJSON(&author); err != nil {
		ctx.JSON(http.StatusBadRequest,
			common.ResponseError(err.Error(), http.StatusBadRequest))
		return
	}
	fmt.Println(author)

	err := ah.authorService.CreateAuthor(&author)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			common.ResponseError(err.Error(), http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, common.ResponseData(author, "Author created",
		http.StatusOK))
}
