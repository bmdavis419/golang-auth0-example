package user

import (
	"github.com/gofiber/fiber/v2"
)

func CreateUserGroup(app *fiber.App, userController *UserController) {
	userGroup := app.Group("/user")

	// middleware to protect routes

	// auth routes
	userGroup.Get("/me", userController.profile)
}
