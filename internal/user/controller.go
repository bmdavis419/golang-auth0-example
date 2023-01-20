package user

import "github.com/gofiber/fiber/v2"

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) profile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "You are logged in",
	})
}
