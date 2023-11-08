package contracts

import (
	"github.com/gofiber/fiber/v2"
)

type MaterialController interface {
	HandlerFindOne(c *fiber.Ctx) error
	HandlerCreate(c *fiber.Ctx) error
	HandlerUpdate(c *fiber.Ctx) error
	HandlerRemove(c *fiber.Ctx) error
}
