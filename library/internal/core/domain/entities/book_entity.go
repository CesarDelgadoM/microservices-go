package entities

import "time"

type Book struct {
	Id       uint      `gorm:"column:id;primaryKey;autoIncrement;index"`
	Title    string    `gorm:"column:title"`
	Price    uint      `gorm:"column:price"`
	Pages    uint      `gorm:"column:pages"`
	CreateAt time.Time `gorm:"column:created_at;autoCreateTime"`
}
