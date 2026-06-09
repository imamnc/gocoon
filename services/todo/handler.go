package todo

import (
	"fmt"
	"strings"

	response "gocoon/core/response"
	"gocoon/database"
	"gocoon/models/entity"
	"gocoon/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetTodo(c *fiber.Ctx) error {
	if c.Query("id") != "" {
		var todo entity.Todo
		if err := database.DB.First(&todo, c.Query("id")).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return response.Success(c, "Successfully to get todo data", nil)
			} else {
				return response.Error(c, fiber.StatusBadRequest, "Failed to get tofo data", err)
			}
		}
		return response.Success(c, "Successfully to get todo data", todo)
	}

	if c.Query("user_id") != "" {
		var todos []entity.Todo
		if err := database.DB.Where("user_id=?", c.Query("user_id")).Find(&todos).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return response.Success(c, "Successfully to get todo data", nil)
			} else {
				return response.Error(c, fiber.StatusBadRequest, "Failed to get todo data", err)
			}
		}
		return response.Success(c, "Successfully to get todo data", todos)
	}

	var todos []entity.Todo
	result := database.DB.Order("title ASC")
	if c.Query("keyword") != "" {
		result.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(c.Query("keyword"))+"%")
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
	// Validate request
	var request CreateTodoRequest
	c.BodyParser(&request)
	if err := request.Validate(); err != nil {
		return response.Validation(c, err)
	}

	// Insert todo data
	var todo entity.Todo
	c.BodyParser(&todo)
	result := database.DB.Create(&todo)
	if result.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to register new todo", result.Error)
	}
	return response.Success(c, "Successfull to create new todo", todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	// Valdate request
	request := UpdateTodoRequest{}
	c.BodyParser(&request)
	if err := request.Validate(); err != nil {
		return response.Validation(c, err)
	}

	// Validate exist
	if !utils.Validation().Exist(database.DB, &entity.Todo{}, "id", request.ID) {
		return response.Error(c, fiber.StatusBadRequest, fmt.Sprintf("Todo with identifier %d does not exists", request.ID), response.InvalidPayload{
			Message: "Todo not found !",
			Value:   request.ID,
			Tag:     "exists",
		})
	}

	var user entity.Todo
	c.BodyParser(&user)
	result := database.DB.Save(&user)
	if result.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update user", result.Error)
	}
	return response.Success(c, "Successfull to update user", user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entity.Todo

	// Validate exist
	if !utils.Validation().Exist(database.DB, &entity.Todo{}, "id", id) {
		return response.Error(c, fiber.StatusBadRequest, fmt.Sprintf("Todo with identifier %v does not exists", id), response.InvalidPayload{
			Message: "Todo not found !",
			Value:   id,
			Tag:     "exists",
		})
	}

	err := database.DB.Where("id=?", id).Delete(&user)
	if err.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to register new user", err)
	}

	return response.Success(c, "Successfully to delete user data", user)
}
