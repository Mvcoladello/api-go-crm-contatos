package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/mvcoladello/api-go-crm-contatos/internal/validators"
)

func SanitizeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "POST" || c.Method() == "PUT" || c.Method() == "PATCH" {
			var body map[string]interface{}

			if err := c.BodyParser(&body); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Formato de dados inválido",
				})
			}

			if nome, ok := body["nome"].(string); ok {
				body["nome"] = validators.SanitizeName(nome)
			}

			if email, ok := body["email"].(string); ok {
				body["email"] = validators.SanitizeEmail(email)
			}

			if cpfCnpj, ok := body["cpf_cnpj"].(string); ok {
				body["cpf_cnpj"] = validators.SanitizeInput(cpfCnpj)
			}

			if telefone, ok := body["telefone"].(string); ok {
				body["telefone"] = validators.SanitizeInput(telefone)
			}

			sanitizedBody, err := json.Marshal(body)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Erro interno do servidor",
				})
			}

			c.Request().SetBody(sanitizedBody)
		}

		return c.Next()
	}
}
