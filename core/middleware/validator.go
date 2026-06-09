package middleware

import (
	"time"

	"gocoon/core/database"
	"gocoon/core/models"
	"gocoon/core/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func DateValidation(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()
	layout := "2006-01-02"
	_, err := time.Parse(layout, dateStr)
	return err == nil
}

func ExistsValidator(fl validator.FieldLevel) bool {
	// Get the current value and model type
	currentValue := fl.Field().String()
	fieldName := fl.Param()                // Get the field name being validated
	modelType := fl.Parent().Type().Name() // Get the model type (e.g., "User", "Product")

	// If field value empty, force return true
	if currentValue == "0" || currentValue == "" {
		return true
	}

	// Use reflection to create a dynamic query
	model := models.Models[modelType]

	// Find the user by ID
	temp := model
	if err := database.DB.Where(fieldName+"=?", currentValue).First(&temp).Error; err != nil {
		return false
	} else {
		return true
	}
}

// Initialize validator
var validate = validator.New()

// ValidateRequest is middleware for validating request bodies
func ValidateRequest(s interface{}) func(c *fiber.Ctx) error {

	validate.RegisterValidation("date", DateValidation)   // Validate date string formatsended
	validate.RegisterValidation("exist", ExistsValidator) // Validate existance of some field

	return func(c *fiber.Ctx) error {
		if err := c.BodyParser(s); err != nil {
			return response.Error(c, fiber.StatusBadRequest, "Request doesn't match excpected payload", err)
		}

		if err := validate.Struct(s); err != nil {
			var errors map[string]response.InvalidPayload = make(map[string]response.InvalidPayload)
			for _, err := range err.(validator.ValidationErrors) {
				errors[err.Field()] = response.InvalidPayload{
					Message: err.Error(),
					Tag:     err.Tag(),
					Param:   err.Param(),
					Value:   err.Value(),
				}
			}
			return response.Error(c, fiber.StatusUnprocessableEntity, "Validation errors", errors)
		}

		return c.Next()
	}
}
