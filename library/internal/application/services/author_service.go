package services

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/application/ports"
	"github.com/CesarDelgadoM/microservices-go/library/internal/domain"
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

func (as AuthorService) FindAuthor(id uint) (domain.Author, error) {
	as.authorRepository.FindById(id)
	return domain.Author{}, nil
}

func (as AuthorService) CreateAuthor(author *domain.Author) error {
	return as.authorRepository.Create(author)
}
