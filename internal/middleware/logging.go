package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		log.Printf("[%s] %s %s - Iniciado",
			c.Method(),
			c.Path(),
			c.IP(),
		)

		err := c.Next()

		duration := time.Since(start)
		statusCode := c.Response().StatusCode()

		log.Printf("[%s] %s %s - Finalizado em %v - Status: %d",
			c.Method(),
			c.Path(),
			c.IP(),
			duration,
			statusCode,
		)

		return err
	}
}

func TracingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		c.Set("X-Request-Start", start.Format(time.RFC3339Nano))

		err := c.Next()

		duration := time.Since(start)
		c.Set("X-Response-Time", duration.String())

		return err
	}
}
