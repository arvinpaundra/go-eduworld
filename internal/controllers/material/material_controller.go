package material

import (
	"github.com/arvinpaundra/go-eduworld/internal/adapters/request"
	"github.com/arvinpaundra/go-eduworld/internal/controllers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/services/material"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type materialController struct {
	materialService material.MaterialService
}

func NewMaterialController(materialService material.MaterialService) contracts.MaterialController {
	return &materialController{materialService: materialService}
}

func (m *materialController) HandlerFindOne(c *fiber.Ctx) error {
	courseId := c.Params("course_id")
	moduleId := c.Params("module_id")
	materialId := c.Params("material_id")

	material, err := m.materialService.FindOne(c.Context(), courseId, moduleId, materialId)

	if err != nil {
		switch err {
		case constants.ErrCourseNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrModuleNotFounnd:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrMaterialNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success get detail material", material))
}

func (m *materialController) HandlerCreate(c *fiber.Ctx) error {
	courseId := c.Params("course_id")
	moduleId := c.Params("module_id")

	var input request.Material

	_ = c.BodyParser(&input)

	if err := helpers.ValidateRequest(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", err))
	}

	err := m.materialService.Save(c.Context(), &input, courseId, moduleId)

	if err != nil {
		switch err {
		case constants.ErrCourseNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrMaterialNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusCreated).JSON(helpers.SuccessCreated("success create material", nil))
}

func (m *materialController) HandlerUpdate(c *fiber.Ctx) error {
	courseId := c.Params("course_id")
	moduleId := c.Params("module_id")
	materialId := c.Params("material_id")

	var input request.Material

	_ = c.BodyParser(&input)

	if err := helpers.ValidateRequest(input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", err))
	}

	err := m.materialService.Update(c.Context(), &input, courseId, moduleId, materialId)

	if err != nil {
		switch err {
		case constants.ErrCourseNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrModuleNotFounnd:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrMaterialNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success update material", nil))
}

func (m *materialController) HandlerRemove(c *fiber.Ctx) error {
	courseId := c.Params("course_id")
	moduleId := c.Params("module_id")
	materialId := c.Params("material_id")

	err := m.materialService.Remove(c.Context(), courseId, moduleId, materialId)

	if err != nil {
		switch err {
		case constants.ErrMaterialNotFound:
			return c.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return c.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success delete material", nil))
}
