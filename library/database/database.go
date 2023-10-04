package database

import (
	"fmt"
	"log"

	"github.com/CesarDelgadoM/microservices-go/library/config"
	"github.com/CesarDelgadoM/microservices-go/library/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type IDatabase interface {
	GetDB() *gorm.DB
}

type Postgres struct {
	*gorm.DB
}

func ConnectPostgres(config *config.PostgresConfig) IDatabase {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=America/Bogota",
		config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the postgres database")
	}

	db.AutoMigrate(&domain.Author{}, &domain.Book{})
	log.Println("Database migration complete")

	log.Println("Connected successfully to the postgres database")
	return &Postgres{
		db,
	}
}

func (postgres Postgres) GetDB() *gorm.DB {
	return postgres.DB
}
