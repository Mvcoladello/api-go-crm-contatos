package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/mvcoladello/api-go-crm-contatos/internal/validators"
)

func ProcessInputMiddleware() fiber.Handler {
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

				if body["nome"].(string) == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Nome é obrigatório",
					})
				}
				if len(body["nome"].(string)) < 2 || len(body["nome"].(string)) > 255 {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Nome deve ter entre 2 e 255 caracteres",
					})
				}
			}

			if email, ok := body["email"].(string); ok {
				body["email"] = validators.SanitizeEmail(email)

				if body["email"].(string) == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Email é obrigatório",
					})
				}
				if !validators.ValidateEmail(body["email"].(string)) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Email inválido",
					})
				}
			}

			if cpfCnpj, ok := body["cpf_cnpj"].(string); ok {
				body["cpf_cnpj"] = validators.SanitizeInput(cpfCnpj)

				if body["cpf_cnpj"].(string) == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "CPF/CNPJ é obrigatório",
					})
				}
				if !validators.ValidateDocument(body["cpf_cnpj"].(string)) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "CPF/CNPJ inválido",
					})
				}
			}

			if telefone, ok := body["telefone"].(string); ok {
				body["telefone"] = validators.SanitizeInput(telefone)

				if body["telefone"].(string) == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Telefone é obrigatório",
					})
				}
				if !validators.ValidateBrazilianPhone(body["telefone"].(string)) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Telefone inválido",
					})
				}
			}

			processedBody, err := json.Marshal(body)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Erro interno do servidor",
				})
			}

			c.Request().SetBody(processedBody)
		}

		return c.Next()
	}
}
