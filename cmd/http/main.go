package main

import (
	"github.com/bmdavis419/golang-auth0-example/internal/user"
	"github.com/bmdavis419/golang-auth0-example/pkg/shutdown"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cleanup, err := run()

	defer cleanup()

	if err != nil {
		panic(err)
	}

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}

func run() (func(), error) {
	app := buildServer()

	// start the server
	go func() {
		app.Listen(":8080")
	}()

	// return a function to close the server and database
	return func() {
		app.Shutdown()
	}, nil
}

func buildServer() *fiber.App {
	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// add health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	// create the user domain
	userController := user.NewUserController()
	user.CreateUserGroup(app, userController)

	return app
}
