package user

import (
	"github.com/bmdavis419/golang-auth0-example/config"
	"github.com/bmdavis419/golang-auth0-example/internal/auth"
	"github.com/gofiber/fiber/v2"
)

func CreateUserGroup(app *fiber.App, userController *UserController, config config.EnvVars) {
	userGroup := app.Group("/user")

	// middleware to protect routes
	authMiddleware := auth.NewAuthMiddleware(config)
	userGroup.Use(authMiddleware.ValidateToken)

	// auth routes
	userGroup.Get("/me", userController.profile)
}
