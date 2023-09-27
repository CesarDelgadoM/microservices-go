package services

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
)

// Application / Business Component
type AuthorService struct {
	authorRepository ports.IAuthorRepository
}

func NewAuthorService(authorRepository ports.IAuthorRepository) ports.IAuthorService {
	return AuthorService{
		authorRepository: authorRepository,
	}
}

func (as AuthorService) FindAuthor(id uint) (entities.Author, error) {
	as.authorRepository.FindById(id)
	return entities.Author{}, nil
}

func (as AuthorService) CreateAuthor(author *entities.Author) error {
	return as.authorRepository.Create(author)
}
