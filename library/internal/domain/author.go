package domain

import "time"

type Author struct {
	Id        uint      `gorm:"column:id;primaryKey;autoIncrement;index"`
	FirstName string    `json:"first_name" gorm:"column:first_name"`
	LastName  string    `json:"last_name" gorm:"column:last_name"`
	Age       uint      `json:"age" gorm:"column:age"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
}
