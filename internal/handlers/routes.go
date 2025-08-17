package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mvcoladello/api-go-crm-contatos/internal/middleware"
	"github.com/mvcoladello/api-go-crm-contatos/internal/services"
)

// SetupRoutes configura todas as rotas da aplicação
func SetupRoutes(app *fiber.App, contactService *services.ContactService) {
	contactHandler := NewContactHandler(contactService)

	api := app.Group("/api/v1")
	api.Use(middleware.LoggingMiddleware())
	api.Use(middleware.TracingMiddleware())

	contacts := api.Group("/contatos")
	contacts.Post("/", middleware.ProcessInputMiddleware(), contactHandler.CreateContact)
	contacts.Get("/", contactHandler.GetContacts)
	contacts.Get("/:id", contactHandler.GetContact)
	contacts.Delete("/:id", contactHandler.DeleteContact)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "API CRM Contatos está funcionando",
		})
	})
}
