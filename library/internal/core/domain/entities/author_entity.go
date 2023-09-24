package entities

import "time"

type Author struct {
	Id        uint      `gorm:"column:id;primaryKey;autoIncrement;index"`
	FirstName string    `gorm:"column:first_name"`
	LastName  string    `gorm:"column:last_name"`
	Age       uint      `gorm:"column:age"`
	BookId    uint      `gorm:"column:book_id;foreignKey:id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}
