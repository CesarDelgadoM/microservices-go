package ports

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"
)

// Primary Port / Driver Port
type AuthorService interface {
	FindAuthor(id uint) (entities.Author, error)
	CreateAuthor(author *entities.Author) error
}

// Secondary Port / Driven Port
type AuthorRepository interface {
	FindById(id uint) (entities.Author, error)
	Create(author *entities.Author) error
}
