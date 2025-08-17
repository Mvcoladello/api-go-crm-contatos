package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/mvcoladello/api-go-crm-contatos/internal/handlers"
	"github.com/mvcoladello/api-go-crm-contatos/internal/middleware"
	"github.com/mvcoladello/api-go-crm-contatos/internal/models"
	"github.com/mvcoladello/api-go-crm-contatos/internal/services"
	"github.com/mvcoladello/api-go-crm-contatos/internal/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// Inicializa o banco de dados
	db, err := initDatabase()
	if err != nil {
		log.Fatal("Falha ao conectar com o banco de dados:", err)
	}

	// Inicializa os services
	contactService := services.NewContactService(db)

	// Inicializa o Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if e, ok := err.(*fiber.Error); ok {
				return utils.SendErrorResponse(c, e.Code, models.ErrCodeInternalServer, e.Message, "")
			}
			return utils.SendInternalServerError(c, "Erro interno do servidor", err.Error())
		},
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}))

	// Rate limiting
	app.Use(middleware.RateLimitByAPI())

	// Swagger documentation
	app.Get("/docs/*", swagger.HandlerDefault)

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":    "OK",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// Configura as rotas
	handlers.SetupRoutes(app, contactService)

	// Inicia o servidor
	log.Println("Servidor iniciando na porta 3000...")
	log.Fatal(app.Listen(":3000"))
}

func initDatabase() (*gorm.DB, error) {
	// Conecta com SQLite
	db, err := gorm.Open(sqlite.Open("crm_contatos.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migração
	err = db.AutoMigrate(&models.Contact{})
	if err != nil {
		return nil, err
	}

	log.Println("Banco de dados conectado e migrado com sucesso!")
	return db, nil
}
