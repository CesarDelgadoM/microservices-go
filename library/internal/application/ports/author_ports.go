package ports

import "github.com/CesarDelgadoM/microservices-go/library/internal/domain"

// Primary Port / Driver Port
type IAuthorService interface {
	FindAuthor(id uint) (domain.Author, error)
	CreateAuthor(author *domain.Author) error
}

// Secondary Port / Driven Port
type IAuthorRepository interface {
	FindById(id uint) (domain.Author, error)
	Create(author *domain.Author) error
}
