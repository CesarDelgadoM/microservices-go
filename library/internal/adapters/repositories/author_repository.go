package repositories

import (
	"github.com/CesarDelgadoM/microservices-go/library/database"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/domain/entities"
	"github.com/CesarDelgadoM/microservices-go/library/internal/core/ports"
)

// Secondary Adapter
type AuthorRepository struct {
	database.IDatabase
}

func NewAuthorRepository(db database.IDatabase) ports.IAuthorRepository {
	return AuthorRepository{
		db,
	}
}

func (ar AuthorRepository) FindById(id uint) (entities.Author, error) {
	return entities.Author{}, nil
}

func (ar AuthorRepository) Create(author *entities.Author) error {
	return ar.GetDB().Create(author).Error
}
