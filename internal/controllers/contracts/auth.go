package contracts

import "github.com/gofiber/fiber/v2"

type AuthController interface {
	HandlerRegister(c *fiber.Ctx) error
	HandlerLogin(c *fiber.Ctx) error
	HandlerLogout(c *fiber.Ctx) error
}
