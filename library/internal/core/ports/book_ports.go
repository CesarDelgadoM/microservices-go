package ports

import "github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"

// Primary Port / Driver Port
type BookService interface {
	FindBook(id uint) (entities.Book, error)
	InsertBook(book *entities.Book) error
}

// Secondary Port / Driven Port
type BookRepository interface {
	FindById(id uint) (entities.Book, error)
	Create(book *entities.Book) error
}
