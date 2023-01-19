package auth

import (
	"fmt"

	"github.com/bmdavis419/golang-auth0-example/config"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	config config.EnvVars
}

func NewAuthMiddleware(config config.EnvVars) *AuthMiddleware {
	return &AuthMiddleware{
		config: config,
	}
}

func (a *AuthMiddleware) ValidateToken(c *fiber.Ctx) error {
	// Set a custom header on all responses:
	c.Set("X-Custom-Header", "Hello, World")

	fmt.Println("hello", a.config.AUTH0_DOMAIN)

	// Go to next middleware:
	return c.Next()
}
