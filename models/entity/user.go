package entity

import (
	"time"

	"gocoon/utils"

	"gorm.io/gorm"
)

// gorm.User definition
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey,autoIncrement"`
	Name      string         `json:"name" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Gender    string         `json:"gender" gorm:"not null"`
	Password  string         `json:"password" gorm:"not null"`
	Birthdate string         `json:"birth_date" gorm:"type:date; not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Todos     []Todo         `json:"todos"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashedPassword, err := utils.Crypt().Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	hashedPassword, err := utils.Crypt().Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = hashedPassword
	return nil
}
