package services

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
	"github.com/gin-gonic/gin"
)

// Application / Business Component
type AuthorService struct {
	authorRepository ports.AuthorRepository
}

func NewAuthorService(authorRepository ports.AuthorRepository) ports.AuthorService {
	return AuthorService{
		authorRepository: authorRepository,
	}
}

func (as AuthorService) FindAuthor(ctx *gin.Context) {

}

func (as AuthorService) InsertAuthor(ctx *gin.Context) {

}
