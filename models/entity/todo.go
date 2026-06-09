package entity

import (
	"time"

	"gorm.io/gorm"
)

// gorm.User definition
type Todo struct {
	ID        uint           `json:"id" gorm:"primaryKey,autoIncrement"`
	UserID    uint           `json:"user_id"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"type:text;unique;not null"`
	Checked   bool           `json:"checked" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	User      User           `json:"user"`
}
