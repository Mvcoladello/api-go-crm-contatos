package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mvcoladello/api-go-crm-contatos/internal/services"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(app *fiber.App, contactService *services.ContactService) {
	// Inicializa os handlers
	contactHandler := NewContactHandler(contactService)

	// Rotas da API
	api := app.Group("/api/v1")

	// Rotas de contatos
	contacts := api.Group("/contatos")
	contacts.Get("/", contactHandler.GetContacts)         // GET /api/v1/contatos
	contacts.Get("/:id", contactHandler.GetContact)       // GET /api/v1/contatos/:id
	contacts.Post("/", contactHandler.CreateContact)      // POST /api/v1/contatos
	contacts.Delete("/:id", contactHandler.DeleteContact) // DELETE /api/v1/contatos/:id

	// Rota de saúde
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API CRM Contatos está funcionando",
		})
	})
}
