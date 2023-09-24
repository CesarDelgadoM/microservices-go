package ports

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"
	"github.com/gin-gonic/gin"
)

// Primary Port / Driver Port
type AuthorService interface {
	FindAuthor(ctx *gin.Context)
	InsertAuthor(ctx *gin.Context)
}

// Secondary Port / Driven Port
type AuthorRepository interface {
	FindById(id uint) entities.Author
	Create(author *entities.Author) error
}
