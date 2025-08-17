package utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mvcoladello/api-go-crm-contatos/internal/models"
)

func SendErrorResponse(c *fiber.Ctx, statusCode int, errorCode, message, details string) error {
	response := models.ErrorResponse{
		Error:     message,
		Code:      errorCode,
		Details:   details,
		Timestamp: time.Now(),
		Path:      c.Path(),
		Method:    c.Method(),
	}

	return c.Status(statusCode).JSON(response)
}

func SendSuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	response := models.SuccessResponse{
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}

	if slice, ok := data.([]interface{}); ok {
		response.Total = len(slice)
	}

	return c.Status(statusCode).JSON(response)
}

func SendSuccessResponseWithTotal(c *fiber.Ctx, statusCode int, message string, data interface{}, total int) error {
	response := models.SuccessResponse{
		Message:   message,
		Data:      data,
		Total:     total,
		Timestamp: time.Now(),
	}

	return c.Status(statusCode).JSON(response)
}

func SendValidationError(c *fiber.Ctx, message string) error {
	return SendErrorResponse(c, fiber.StatusBadRequest, models.ErrCodeValidation, message, "")
}

func SendNotFoundError(c *fiber.Ctx, message string) error {
	return SendErrorResponse(c, fiber.StatusNotFound, models.ErrCodeNotFound, message, "")
}

func SendConflictError(c *fiber.Ctx, message string) error {
	return SendErrorResponse(c, fiber.StatusConflict, models.ErrCodeConflict, message, "")
}

func SendInternalServerError(c *fiber.Ctx, message, details string) error {
	return SendErrorResponse(c, fiber.StatusInternalServerError, models.ErrCodeInternalServer, message, details)
}

func SendBadRequestError(c *fiber.Ctx, message string) error {
	return SendErrorResponse(c, fiber.StatusBadRequest, models.ErrCodeBadRequest, message, "")
}
