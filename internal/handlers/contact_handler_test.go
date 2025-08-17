package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mvcoladello/api-go-crm-contatos/internal/models"
	"github.com/mvcoladello/api-go-crm-contatos/internal/services"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Contact{})
	return db
}

func setupTestApp() (*fiber.App, *services.ContactService) {
	db := setupTestDB()
	contactService := services.NewContactService(db)

	app := fiber.New()
	SetupRoutes(app, contactService)

	return app, contactService
}

func TestGetContacts(t *testing.T) {
	app, _ := setupTestApp()

	req := httptest.NewRequest("GET", "/api/v1/contatos", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Status esperado 200, obtido %d", resp.StatusCode)
	}
}

func TestCreateContact(t *testing.T) {
	app, _ := setupTestApp()

	contact := models.Contact{
		Nome:     "João Silva",
		Email:    "joao@test.com",
		CPFCNPJ:  "11144477735", // CPF válido
		Telefone: "11999999999",
	}

	body, _ := json.Marshal(contact)
	req := httptest.NewRequest("POST", "/api/v1/contatos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Status esperado 201, obtido %d", resp.StatusCode)
	}
}

func TestCreateContactInvalidCPF(t *testing.T) {
	app, _ := setupTestApp()

	contact := models.Contact{
		Nome:     "João Silva",
		Email:    "joao@test.com",
		CPFCNPJ:  "12345678901", // CPF inválido
		Telefone: "11999999999",
	}

	body, _ := json.Marshal(contact)
	req := httptest.NewRequest("POST", "/api/v1/contatos", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Status esperado 400, obtido %d", resp.StatusCode)
	}
}

func TestGetContactByID(t *testing.T) {
	app, service := setupTestApp()

	// Criar um contato primeiro
	contact := &models.Contact{
		Nome:     "João Silva",
		Email:    "joao@test.com",
		CPFCNPJ:  "11144477735",
		Telefone: "11999999999",
	}

	err := service.CreateContact(contact)
	if err != nil {
		t.Fatalf("Erro ao criar contato: %v", err)
	}

	// Buscar o contato
	req := httptest.NewRequest("GET", "/api/v1/contatos/"+contact.ID.String(), nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Status esperado 200, obtido %d", resp.StatusCode)
	}
}

func TestGetContactByInvalidID(t *testing.T) {
	app, _ := setupTestApp()

	req := httptest.NewRequest("GET", "/api/v1/contatos/invalid-id", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("Status esperado 400, obtido %d", resp.StatusCode)
	}
}

func TestGetContactNotFound(t *testing.T) {
	app, _ := setupTestApp()

	randomID := uuid.New()
	req := httptest.NewRequest("GET", "/api/v1/contatos/"+randomID.String(), nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Status esperado 404, obtido %d", resp.StatusCode)
	}
}

func TestDeleteContact(t *testing.T) {
	app, service := setupTestApp()

	// Criar um contato primeiro
	contact := &models.Contact{
		Nome:     "João Silva",
		Email:    "joao@test.com",
		CPFCNPJ:  "11144477735",
		Telefone: "11999999999",
	}

	err := service.CreateContact(contact)
	if err != nil {
		t.Fatalf("Erro ao criar contato: %v", err)
	}

	// Deletar o contato
	req := httptest.NewRequest("DELETE", "/api/v1/contatos/"+contact.ID.String(), nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Status esperado 200, obtido %d", resp.StatusCode)
	}
}

func TestDeleteContactNotFound(t *testing.T) {
	app, _ := setupTestApp()

	randomID := uuid.New()
	req := httptest.NewRequest("DELETE", "/api/v1/contatos/"+randomID.String(), nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("Status esperado 404, obtido %d", resp.StatusCode)
	}
}
