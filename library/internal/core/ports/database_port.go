package ports

import "gorm.io/gorm"

type Database interface {
	GetDB() *gorm.DB
}
