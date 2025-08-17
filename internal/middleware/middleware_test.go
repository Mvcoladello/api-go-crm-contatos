package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestProcessInputMiddleware(t *testing.T) {
	app := fiber.New()

	app.Post("/test", ProcessInputMiddleware(), func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	tests := []struct {
		name          string
		body          string
		expectedCode  int
		expectedError string
	}{
		{
			name: "Valid input",
			body: `{
				"nome": "João Silva",
				"email": "joao@test.com",
				"cpf_cnpj": "11144477735",
				"telefone": "11999999999"
			}`,
			expectedCode: 200,
		},
		{
			name: "Empty name",
			body: `{
				"nome": "",
				"email": "joao@test.com",
				"cpf_cnpj": "11144477735",
				"telefone": "11999999999"
			}`,
			expectedCode:  400,
			expectedError: "Nome é obrigatório",
		},
		{
			name: "Invalid email",
			body: `{
				"nome": "João Silva",
				"email": "email-invalido",
				"cpf_cnpj": "11144477735",
				"telefone": "11999999999"
			}`,
			expectedCode:  400,
			expectedError: "Email inválido",
		},
		{
			name: "Invalid CPF",
			body: `{
				"nome": "João Silva",
				"email": "joao@test.com",
				"cpf_cnpj": "12345678901",
				"telefone": "11999999999"
			}`,
			expectedCode:  400,
			expectedError: "CPF/CNPJ inválido",
		},
		{
			name: "Invalid phone",
			body: `{
				"nome": "João Silva",
				"email": "joao@test.com",
				"cpf_cnpj": "11144477735",
				"telefone": "123"
			}`,
			expectedCode:  400,
			expectedError: "Telefone inválido",
		},
		{
			name: "Sanitization test",
			body: `{
				"nome": "  João Silva  ",
				"email": "  JOAO@TEST.COM  ",
				"cpf_cnpj": "111.444.777-35",
				"telefone": "(11) 99999-9999"
			}`,
			expectedCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/test", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Erro na requisição: %v", err)
			}

			if resp.StatusCode != tt.expectedCode {
				t.Errorf("Status esperado %d, obtido %d", tt.expectedCode, resp.StatusCode)
			}
		})
	}
}

func TestLoggingMiddleware(t *testing.T) {
	app := fiber.New()

	app.Use(LoggingMiddleware())
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status esperado 200, obtido %d", resp.StatusCode)
	}
}

func TestTracingMiddleware(t *testing.T) {
	app := fiber.New()

	app.Use(TracingMiddleware())
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	req := httptest.NewRequest("GET", "/test", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Erro na requisição: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status esperado 200, obtido %d", resp.StatusCode)
	}

	responseTime := resp.Header.Get("X-Response-Time")
	if responseTime == "" {
		t.Error("Header X-Response-Time não foi definido")
	}
}
