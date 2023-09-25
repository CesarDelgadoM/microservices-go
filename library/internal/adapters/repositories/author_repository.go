package repositories

import (
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
)

// Secondary Adapter
type AuthorRepository struct {
	db ports.Database
}

func NewAuthorRepository(db ports.Database) ports.AuthorRepository {
	return AuthorRepository{
		db: db,
	}
}

func (ar AuthorRepository) FindById(id uint) (entities.Author, error) {
	return entities.Author{}, nil
}

func (ar AuthorRepository) Create(author *entities.Author) error {
	return nil
}
