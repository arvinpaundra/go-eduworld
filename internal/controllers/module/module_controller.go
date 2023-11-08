package module

import (
	"github.com/arvinpaundra/go-eduworld/internal/adapters/request"
	"github.com/arvinpaundra/go-eduworld/internal/controllers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/services/module"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type moduleController struct {
	moduleService module.ModuleService
}

func NewModuleController(moduleService module.ModuleService) contracts.ModuleController {
	return &moduleController{moduleService: moduleService}
}

func (m *moduleController) HandlerCreate(c *fiber.Ctx) error {
	courseId := c.Params("course_id")

	var input request.Module

	_ = c.BodyParser(&input)

	if err := helpers.ValidateRequest(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", err))
	}

	err := m.moduleService.Save(c.Context(), &input, courseId)

	if err != nil {
		switch err {
		case constants.ErrCourseNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusCreated).JSON(helpers.SuccessCreated("success create module", nil))
}

func (m *moduleController) HandlerUpdate(c *fiber.Ctx) error {
	courseId := c.Params("course_id")
	moduleId := c.Params("module_id")

	var input request.Module

	_ = c.BodyParser(&input)

	if err := helpers.ValidateRequest(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", err))
	}

	err := m.moduleService.Update(c.Context(), &input, courseId, moduleId)

	if err != nil {
		switch err {
		case constants.ErrCourseNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrModuleNotFounnd:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success update module", nil))
}

func (m *moduleController) HandlerRemove(c *fiber.Ctx) error {
	courseId := c.Params("course_id")
	moduleId := c.Params("module_id")

	err := m.moduleService.Remove(c.Context(), courseId, moduleId)

	if err != nil {
		switch err {
		case constants.ErrModuleNotFounnd:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success delete module", nil))
}
