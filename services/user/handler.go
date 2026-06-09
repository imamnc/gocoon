package user

import (
	"fmt"
	"strings"

	"gocoon/core/database"
	"gocoon/core/models/entity"
	"gocoon/core/response"
	"gocoon/core/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetUser(c *fiber.Ctx) error {
	if c.Query("id") != "" {
		var user entity.User
		if err := database.DB.Preload("Todos").First(&user, c.Query("id")).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return response.Success(c, "Successfully to get data user", nil)
			} else {
				return response.Error(c, fiber.StatusBadRequest, "Failed to get user data", err)
			}
		}
		return response.Success(c, "Successfully to get data user", user)
	}

	var users []entity.User
	result := database.DB.Preload("Todos").Order("name ASC")
	if c.Query("keyword") != "" {
		result.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(c.Query("keyword"))+"%")
	}
	result.Find(&users)
	if err := result.Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to get users data", err)
		}
	}

	return response.Success(c, "Successfully to get the users data", users)
}

func RegisterUser(c *fiber.Ctx) error {
	// Validate request
	var userRequest CreateUserRequest
	c.BodyParser(&userRequest)
	if err := userRequest.Validate(); err != nil {
		return response.Validation(c, err)
	}

	// Validate unique
	if !utils.Validation().Unique(database.DB, &entity.User{}, "email", userRequest.Email) {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Failed to register new user", response.InvalidPayload{
			Message: "Email already used !",
			Value:   userRequest.Email,
			Tag:     "unique",
		})
	}

	// Insert user data
	var user entity.User
	c.BodyParser(&user)
	result := database.DB.Create(&user)
	if result.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to register new user", result.Error)
	}
	return response.Success(c, "Successfull to create new user", user)
}

func UpdateUser(c *fiber.Ctx) error {
	// Valdate request
	userRequest := UpdateUserRequest{}
	c.BodyParser(&userRequest)
	if err := userRequest.Validate(); err != nil {
		return response.Validation(c, err)
	}

	// Validate exist
	if !utils.Validation().Exist(database.DB, &entity.User{}, "id", userRequest.ID) {
		return response.Error(c, fiber.StatusBadRequest, fmt.Sprintf("User with identifier %d does not exists", userRequest.ID), response.InvalidPayload{
			Message: "User not found !",
			Value:   userRequest.ID,
			Tag:     "exists",
		})
	}

	// Validate unique
	if !utils.Validation().Unique(database.DB, &entity.User{}, "email", userRequest.Email, userRequest.ID) {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Failed to update user", response.InvalidPayload{
			Message: "Email already used !",
			Value:   userRequest.Email,
			Tag:     "unique",
		})
	}

	var user entity.User
	c.BodyParser(&user)
	result := database.DB.Save(&user)
	if result.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update user", result.Error)
	}
	return response.Success(c, "Successfull to update user", user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user entity.User

	// Validate exist
	if !utils.Validation().Exist(database.DB, &entity.User{}, "id", id) {
		return response.Error(c, fiber.StatusBadRequest, fmt.Sprintf("User with identifier %v does not exists", id), response.InvalidPayload{
			Message: "User not found !",
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
