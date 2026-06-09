package auth

import (
	"fmt"

	"gocoon/core/database"
	"gocoon/core/models/entity"
	"gocoon/core/response"
	"gocoon/core/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	var request LoginRequest
	c.BodyParser(&request)
	if err := request.Validate(); err != nil {
		return response.Validation(c, err)
	}

	// Get user by email
	var user entity.User
	query := database.DB.Where("email=?", request.Email).First(&user)
	if query.Error != nil {
		if query.Error == gorm.ErrRecordNotFound {
			return response.Error(c, fiber.StatusUnprocessableEntity, fmt.Sprintf("User with email %v not found !", request.Email), query.Error)
		} else {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to login, code: x001", query.Error)
		}
	}

	// Validate password
	errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if errPassword != nil {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Incorrect Password !", errPassword)
	}

	// Create JWT Token
	token, err := utils.Jwt().CreateToken(int(user.ID))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to login, code: x002", err)
	}

	// Response
	return response.Success(c, "Successfully to create access token", fiber.Map{
		"token": token,
	})
}
