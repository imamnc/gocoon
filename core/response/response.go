package response

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

// Define invalid payload response
type InvalidPayload struct {
	Message string `json:"message"`
	Tag     string `json:"tag,omitempty"`
	Value   any    `json:"value"`
	Param   any    `json:"param,omitempty"`
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(c *fiber.Ctx, statusCode int, message string, err interface{}) error {
	switch e := err.(type) {
	case error:
		return c.Status(statusCode).JSON(ErrorResponse{
			Status:  "error",
			Message: message,
			Error:   e.Error(),
		})
	default:
		return c.Status(statusCode).JSON(ErrorResponse{
			Status:  "error",
			Message: message,
			Error:   err,
		})
	}
}

// Response validation
func Validation(c *fiber.Ctx, result interface{}) error {
	if result != nil {
		return Error(c, 422, "Validation Error !", result)
	}
	return nil
}
