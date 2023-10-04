package repositories

import (
	"github.com/CesarDelgadoM/microservices-go/library/database"
	"github.com/CesarDelgadoM/microservices-go/library/internal/application/ports"
	"github.com/CesarDelgadoM/microservices-go/library/internal/domain"
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

func (ar AuthorRepository) FindById(id uint) (domain.Author, error) {
	return domain.Author{}, nil
}

func (ar AuthorRepository) Create(author *domain.Author) error {
	return ar.GetDB().Create(author).Error
}
