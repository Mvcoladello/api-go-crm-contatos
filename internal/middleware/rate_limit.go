package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/mvcoladello/api-go-crm-contatos/internal/models"
	"github.com/mvcoladello/api-go-crm-contatos/internal/utils"
	"golang.org/x/time/rate"
)

// RateLimitConfig configuração do rate limiter
type RateLimitConfig struct {
	// Máximo de requisições por segundo
	Rate rate.Limit
	// Burst máximo de requisições em rajada
	Burst int
	// Duração da janela de tempo
	Window time.Duration
	// Função para extrair chave de identificação (padrão: IP)
	KeyGenerator func(c *fiber.Ctx) string
	// Mensagem de erro personalizada
	Message string
	// Skip permite pular o rate limiting para certas condições
	Skip func(c *fiber.Ctx) bool
}

// Cliente representa um cliente com rate limiter
type rateLimitClient struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Store para armazenar os rate limiters por cliente
type store struct {
	clients map[string]*rateLimitClient
	mu      sync.RWMutex
}

// Instância global do store
var rateLimitStore = &store{
	clients: make(map[string]*rateLimitClient),
}

// RateLimiter cria um middleware de rate limiting
func RateLimiter(config ...RateLimitConfig) fiber.Handler {
	// Configuração padrão
	cfg := RateLimitConfig{
		Rate:   rate.Limit(10), // 10 requisições por segundo
		Burst:  20,             // Burst de até 20 requisições
		Window: time.Minute,    // Janela de 1 minuto
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		Message: "Rate limit exceeded. Too many requests.",
		Skip: func(c *fiber.Ctx) bool {
			return false
		},
	}

	// Aplica configuração personalizada se fornecida
	if len(config) > 0 {
		if config[0].Rate > 0 {
			cfg.Rate = config[0].Rate
		}
		if config[0].Burst > 0 {
			cfg.Burst = config[0].Burst
		}
		if config[0].Window > 0 {
			cfg.Window = config[0].Window
		}
		if config[0].KeyGenerator != nil {
			cfg.KeyGenerator = config[0].KeyGenerator
		}
		if config[0].Message != "" {
			cfg.Message = config[0].Message
		}
		if config[0].Skip != nil {
			cfg.Skip = config[0].Skip
		}
	}

	// Limpa clientes antigos periodicamente
	go cleanupRoutine(cfg.Window)

	return func(c *fiber.Ctx) error {
		// Skip se a condição for atendida
		if cfg.Skip(c) {
			return c.Next()
		}

		key := cfg.KeyGenerator(c)
		if key == "" {
			return c.Next()
		}

		// Obtém ou cria o rate limiter para este cliente
		limiter := getRateLimiter(key, cfg.Rate, cfg.Burst)

		// Verifica se a requisição é permitida
		if !limiter.Allow() {
			return utils.SendErrorResponse(
				c,
				fiber.StatusTooManyRequests,
				models.ErrCodeValidation,
				cfg.Message,
				"Rate limit exceeded",
			)
		}

		return c.Next()
	}
}

// getRateLimiter obtém ou cria um rate limiter para uma chave específica
func getRateLimiter(key string, rateLimit rate.Limit, burst int) *rate.Limiter {
	rateLimitStore.mu.Lock()
	defer rateLimitStore.mu.Unlock()

	client, exists := rateLimitStore.clients[key]
	if !exists {
		limiter := rate.NewLimiter(rateLimit, burst)
		rateLimitStore.clients[key] = &rateLimitClient{
			limiter:  limiter,
			lastSeen: time.Now(),
		}
		return limiter
	}

	client.lastSeen = time.Now()
	return client.limiter
}

// cleanupRoutine remove clientes inativos periodicamente
func cleanupRoutine(window time.Duration) {
	ticker := time.NewTicker(window)
	defer ticker.Stop()

	for range ticker.C {
		cleanup(window)
	}
}

// cleanup remove clientes que não fizeram requisições há um tempo
func cleanup(window time.Duration) {
	rateLimitStore.mu.Lock()
	defer rateLimitStore.mu.Unlock()

	cutoff := time.Now().Add(-window * 3) // Remove clientes inativos há 3x a janela

	for key, client := range rateLimitStore.clients {
		if client.lastSeen.Before(cutoff) {
			delete(rateLimitStore.clients, key)
		}
	}
}

// GetRateLimitStats retorna estatísticas do rate limiter
func GetRateLimitStats() map[string]interface{} {
	rateLimitStore.mu.RLock()
	defer rateLimitStore.mu.RUnlock()

	return map[string]interface{}{
		"active_clients": len(rateLimitStore.clients),
		"timestamp":      time.Now(),
	}
}

// RateLimitByAPI cria configurações específicas por endpoint
func RateLimitByAPI() fiber.Handler {
	return RateLimiter(RateLimitConfig{
		Rate:  rate.Limit(100), // 100 req/s para APIs normais
		Burst: 200,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		Skip: func(c *fiber.Ctx) bool {
			// Skip para endpoints de health check
			return c.Path() == "/health"
		},
	})
}

// RateLimitStrict configuração mais restritiva para endpoints sensíveis
func RateLimitStrict() fiber.Handler {
	return RateLimiter(RateLimitConfig{
		Rate:  rate.Limit(5), // 5 req/s
		Burst: 10,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		Message: "Rate limit exceeded for sensitive operation",
	})
}
