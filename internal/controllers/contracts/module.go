package contracts

import "github.com/gofiber/fiber/v2"

type ModuleController interface {
	HandlerCreate(c *fiber.Ctx) error
	HandlerUpdate(c *fiber.Ctx) error
	HandlerRemove(c *fiber.Ctx) error
}
