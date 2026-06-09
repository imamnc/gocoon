package todo

import (
	"gocoon/core/utils"

	"github.com/thedevsaddam/govalidator"
)

type CreateTodoRequest struct {
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Checked bool   `json:"checked"`
}

func (u *CreateTodoRequest) Validate() interface{} {
	err := utils.Validation().Validate(u, govalidator.MapData{
		"user_id": []string{"required"},
		"title":   []string{"required"},
		"content": []string{"required"},
		"checked": []string{"bool"},
	})
	return err
}

type UpdateTodoRequest struct {
	ID      int    `json:"id"`
	UserID  int    `json:"user_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Checked bool   `json:"checked"`
}

func (u *UpdateTodoRequest) Validate() interface{} {
	err := utils.Validation().Validate(u, govalidator.MapData{
		"id":      []string{"required"},
		"user_id": []string{"required"},
		"title":   []string{"required"},
		"content": []string{"required"},
		"checked": []string{"bool"},
	})
	return err
}
