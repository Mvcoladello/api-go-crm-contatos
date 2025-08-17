package middleware

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/mvcoladello/api-go-crm-contatos/internal/validators"
)

func ValidationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Method() == "POST" || c.Method() == "PUT" || c.Method() == "PATCH" {
			var body map[string]interface{}

			if err := c.BodyParser(&body); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Formato de dados inválido",
				})
			}

			if nome, ok := body["nome"].(string); ok {
				if nome == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Nome é obrigatório",
					})
				}
				if len(nome) < 2 || len(nome) > 255 {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Nome deve ter entre 2 e 255 caracteres",
					})
				}
			}

			if email, ok := body["email"].(string); ok {
				if email == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Email é obrigatório",
					})
				}
				if !validators.ValidateEmail(email) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Email inválido",
					})
				}
			}

			if cpfCnpj, ok := body["cpf_cnpj"].(string); ok {
				if cpfCnpj == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "CPF/CNPJ é obrigatório",
					})
				}
				if !validators.ValidateDocument(cpfCnpj) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "CPF/CNPJ inválido",
					})
				}
			}

			if telefone, ok := body["telefone"].(string); ok {
				if telefone == "" {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Telefone é obrigatório",
					})
				}
				if !validators.ValidateBrazilianPhone(telefone) {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Telefone inválido",
					})
				}
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
