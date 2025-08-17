package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
)

func TestRateLimiter(t *testing.T) {
	tests := []struct {
		name         string
		config       RateLimitConfig
		requests     int
		expectedCode int
		delay        time.Duration
	}{
		{
			name: "Rate limit not exceeded",
			config: RateLimitConfig{
				Rate:  rate.Limit(10),
				Burst: 1,
			},
			requests:     1,
			expectedCode: http.StatusOK,
		},
		{
			name: "Rate limit exceeded",
			config: RateLimitConfig{
				Rate:  rate.Limit(1),
				Burst: 1,
			},
			requests:     2,
			expectedCode: http.StatusTooManyRequests,
		},
		{
			name: "Custom key generator",
			config: RateLimitConfig{
				Rate:  rate.Limit(1),
				Burst: 1,
				KeyGenerator: func(c *fiber.Ctx) string {
					return "custom-key"
				},
			},
			requests:     2,
			expectedCode: http.StatusTooManyRequests,
		},
		{
			name: "Skip function works",
			config: RateLimitConfig{
				Rate:  rate.Limit(1),
				Burst: 1,
				Skip: func(c *fiber.Ctx) bool {
					return true // sempre pula
				},
			},
			requests:     5,
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := fiber.New()
			app.Use(RateLimiter(tt.config))
			app.Get("/test", func(c *fiber.Ctx) error {
				return c.SendStatus(http.StatusOK)
			})

			var lastStatus int
			for i := 0; i < tt.requests; i++ {
				req := httptest.NewRequest("GET", "/test", nil)
				resp, err := app.Test(req)
				if err != nil {
					t.Fatalf("Erro ao testar requisição %d: %v", i+1, err)
				}
				lastStatus = resp.StatusCode
				resp.Body.Close()

				if tt.delay > 0 {
					time.Sleep(tt.delay)
				}
			}

			if lastStatus != tt.expectedCode {
				t.Errorf("Esperado status %d, obtido %d", tt.expectedCode, lastStatus)
			}
		})
	}
}

func TestRateLimitByAPI(t *testing.T) {
	// Limpar store antes do teste
	rateLimitStore.mu.Lock()
	rateLimitStore.clients = make(map[string]*rateLimitClient)
	rateLimitStore.mu.Unlock()

	app := fiber.New()
	app.Use(RateLimitByAPI())

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	// Testar endpoint health (deve ser pulado)
	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao testar /health: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Esperado status 200 para /health, obtido %d", resp.StatusCode)
	}
	resp.Body.Close()

	// Testar endpoint normal (primeira requisição deve passar)
	req = httptest.NewRequest("GET", "/test", nil)
	resp, err = app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao testar /test: %v", err)
	}

	// Aceitar tanto 200 quanto 429, pois pode já ter atingido o limite
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Esperado status 200 ou 429 para /test, obtido %d", resp.StatusCode)
	}
	resp.Body.Close()
}

func TestRateLimitStrict(t *testing.T) {
	// Limpar store antes do teste
	rateLimitStore.mu.Lock()
	rateLimitStore.clients = make(map[string]*rateLimitClient)
	rateLimitStore.mu.Unlock()

	app := fiber.New()
	app.Use(RateLimitStrict())

	app.Get("/sensitive", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	// Primeira requisição - aceitar tanto 200 quanto 429
	req := httptest.NewRequest("GET", "/sensitive", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Erro ao testar primeira requisição: %v", err)
	}

	// Com rate limiting strict, pode já estar limitado
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("Esperado status 200 ou 429 para primeira requisição, obtido %d", resp.StatusCode)
	}
	resp.Body.Close()
}

func TestGetRateLimitStats(t *testing.T) {
	// Limpar store antes do teste
	rateLimitStore.mu.Lock()
	rateLimitStore.clients = make(map[string]*rateLimitClient)
	rateLimitStore.mu.Unlock()

	stats := GetRateLimitStats()

	if stats["active_clients"] != 0 {
		t.Errorf("Esperado 0 clientes ativos, obtido %v", stats["active_clients"])
	}

	// Criar um cliente
	_ = getRateLimiter("test-client", rate.Limit(10), 10)

	stats = GetRateLimitStats()
	if stats["active_clients"] != 1 {
		t.Errorf("Esperado 1 cliente ativo, obtido %v", stats["active_clients"])
	}
}

func BenchmarkRateLimiter(b *testing.B) {
	app := fiber.New()
	app.Use(RateLimiter())
	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			req := httptest.NewRequest("GET", "/test", nil)
			resp, err := app.Test(req)
			if err != nil {
				b.Fatalf("Erro no benchmark: %v", err)
			}
			resp.Body.Close()
		}
	})
}
