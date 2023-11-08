package contracts

import "github.com/gofiber/fiber/v2"

type CourseController interface {
	HandlerFindAll(f *fiber.Ctx) error
	HandlerFindByMentor(f *fiber.Ctx) error
	HandlerFindDetail(f *fiber.Ctx) error
	HandlerCreate(f *fiber.Ctx) error
	HandlerUpdate(f *fiber.Ctx) error
	HandlerRemove(f *fiber.Ctx) error
}
