package ports

import "github.com/CesarDelgadoM/microservices-go/library/internal/domain"

// Primary Port / Driver Port
type BookService interface {
	FindBook(id uint) (domain.Book, error)
	InsertBook(book *domain.Book) error
}

// Secondary Port / Driven Port
type BookRepository interface {
	FindById(id uint) (domain.Book, error)
	Create(book *domain.Book) error
}
