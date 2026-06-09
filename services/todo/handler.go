package todo

import (
	"errors"
	"fmt"
	"strings"

	"gocoon/core/database"
	"gocoon/core/models/entity"
	"gocoon/core/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetTodo(c *fiber.Ctx) error {
	if c.Query("id") != "" {
		var todo entity.Todo
		if err := database.DB.First(&todo, c.Query("id")).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return response.Success(c, "Successfully to get todo data", nil)
			}
			return response.Error(c, fiber.StatusBadRequest, "Failed to get todo data", err)
		}
		return response.Success(c, "Successfully to get todo data", todo)
	}

	if c.Query("user_id") != "" {
		var todos []entity.Todo
		if err := database.DB.Where("user_id=?", c.Query("user_id")).Find(&todos).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return response.Success(c, "Successfully to get todo data", nil)
			}
			return response.Error(c, fiber.StatusBadRequest, "Failed to get todo data", err)
		}
		return response.Success(c, "Successfully to get todo data", todos)
	}

	var todos []entity.Todo
	result := database.DB.Order("title ASC")
	if c.Query("keyword") != "" {
		result = result.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(c.Query("keyword"))+"%")
	}
	result.Find(&todos)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to get todos data", err)
		}
	}

	return response.Success(c, "Successfully to get the todos data", todos)
}

func CreateTodo(c *fiber.Ctx) error {
	var request CreateTodoRequest
	if err := c.BodyParser(&request); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err)
	}

	if err := request.Validate(); err != nil {
		return response.Validation(c, err)
	}

	todo := entity.Todo{
		UserID:  uint(request.UserID),
		Title:   request.Title,
		Content: request.Content,
		Checked: request.Checked,
	}

	result := database.DB.Create(&todo)
	if result.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create new todo", result.Error)
	}

	return response.Success(c, "Successfully created todo", todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	var request UpdateTodoRequest
	if err := c.BodyParser(&request); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request payload", err)
	}

	if err := request.Validate(); err != nil {
		return response.Validation(c, err)
	}

	if !todoExists(request.ID) {
		return response.Error(c, fiber.StatusBadRequest, fmt.Sprintf("Todo with identifier %d does not exist", request.ID), response.InvalidPayload{
			Message: "Todo not found!",
			Value:   request.ID,
			Tag:     "exists",
		})
	}

	todo := entity.Todo{
		ID:      uint(request.ID),
		UserID:  uint(request.UserID),
		Title:   request.Title,
		Content: request.Content,
		Checked: request.Checked,
	}

	result := database.DB.Save(&todo)
	if result.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update todo", result.Error)
	}
	return response.Success(c, "Successfully updated todo", todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo entity.Todo

	if !todoExists(id) {
		return response.Error(c, fiber.StatusBadRequest, fmt.Sprintf("Todo with identifier %v does not exist", id), response.InvalidPayload{
			Message: "Todo not found!",
			Value:   id,
			Tag:     "exists",
		})
	}

	err := database.DB.Where("id=?", id).Delete(&todo)
	if err.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete todo", err)
	}

	return response.Success(c, "Successfully deleted todo", todo)
}

func todoExists(id interface{}) bool {
	var todo entity.Todo
	err := database.DB.First(&todo, id).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
