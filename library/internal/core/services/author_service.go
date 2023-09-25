package services

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
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

func (as AuthorService) FindAuthor(id uint) (entities.Author, error) {
	as.authorRepository.FindById(id)
	return entities.Author{}, nil
}

func (as AuthorService) CreateAuthor(author *entities.Author) error {
	as.authorRepository.Create(author)
	return nil
}
