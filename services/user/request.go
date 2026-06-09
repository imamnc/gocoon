package user

import (
	"gocoon/utils"

	"github.com/thedevsaddam/govalidator"
)

/* CREATE USER REQUEST */
type CreateUserRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Password  string `json:"password"`
	Birthdate string `json:"birth_date"`
}

func (u *CreateUserRequest) Validate() interface{} {
	err := utils.Validation().Validate(u, govalidator.MapData{
		"name":       []string{"required"},
		"email":      []string{"required", "email"},
		"gender":     []string{"required", "in:male,female"},
		"password":   []string{"required"},
		"birth_date": []string{"required", "date:yyyy-mm-dd"},
	})
	return err
}

/* UPDATE USER REQUEST */
type UpdateUserRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Password  string `json:"password"`
	Birthdate string `json:"birth_date"`
}

func (u *UpdateUserRequest) Validate() interface{} {
	err := utils.Validation().Validate(u, govalidator.MapData{
		"id":         []string{"required"},
		"name":       []string{"required"},
		"email":      []string{"required", "email"},
		"gender":     []string{"required", "in:male,female"},
		"password":   []string{"required"},
		"birth_date": []string{"required", "date:yyyy-mm-dd"},
	})
	return err
}
