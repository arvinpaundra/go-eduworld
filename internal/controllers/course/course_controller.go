package course

import (
	"github.com/arvinpaundra/go-eduworld/internal/adapters/request"
	"github.com/arvinpaundra/go-eduworld/internal/controllers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/services/course"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/helpers"
	"github.com/gofiber/fiber/v2"
)

type courseController struct {
	courseService course.CourseService
}

func NewCourseController(courseService course.CourseService) contracts.CourseController {
	return &courseController{courseService: courseService}
}

func (c *courseController) HandlerFindAll(f *fiber.Ctx) error {
	query := request.QueryFindCourses{
		Keyword:  f.Query("keyword"),
		Category: f.Query("category"),
		Page:     f.QueryInt("page") - 1,
	}

	results, pagination, err := c.courseService.Find(f.Context(), &query)

	if err != nil {
		switch err {
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success find all courses", results, helpers.Pagination{OffsetPagination: pagination}))
}

func (c *courseController) HandlerFindByMentor(f *fiber.Ctx) error {
	mentorId := f.Params("mentor_id")

	query := request.QueryFindCourses{
		Keyword:  f.Query("keyword"),
		Category: f.Query("category"),
		Page:     f.QueryInt("page") - 1,
	}

	results, pagination, err := c.courseService.FindByMentor(f.Context(), mentorId, &query)

	if err != nil {
		switch err {
		case constants.ErrUserNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success get all mentor courses", results, helpers.Pagination{OffsetPagination: pagination}))
}

func (c *courseController) HandlerFindDetail(f *fiber.Ctx) error {
	courseId := f.Params("course_id")

	result, err := c.courseService.DetailCourse(f.Context(), courseId)

	if err != nil {
		switch err {
		case constants.ErrCourseNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success get detail course", result))
}

func (c *courseController) HandlerCreate(f *fiber.Ctx) error {
	var input request.Course

	_ = f.BodyParser(&input)

	if err := helpers.ValidateRequest(input); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", err))
	}

	err := c.courseService.Save(f.Context(), &input)

	if err != nil {
		switch err {
		case constants.ErrCategoryNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrInterestNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusCreated).JSON(helpers.SuccessCreated("success create course", nil))
}

func (c *courseController) HandlerUpdate(f *fiber.Ctx) error {
	var input request.Course

	_ = f.BodyParser(&input)

	if err := helpers.ValidateRequest(input); err != nil {
		return f.Status(fiber.StatusBadRequest).JSON(helpers.BadRequest("invalid request body", err))
	}

	courseId := f.Params("course_id")

	err := c.courseService.Update(f.Context(), &input, courseId)

	if err != nil {
		switch err {
		case constants.ErrCourseNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrCategoryNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrUserNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		case constants.ErrInterestNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success update course", nil))
}

func (c *courseController) HandlerRemove(f *fiber.Ctx) error {
	courseId := f.Params("course_id")

	err := c.courseService.Remove(f.Context(), courseId)

	if err != nil {
		switch err {
		case constants.ErrCourseNotFound:
			return f.Status(fiber.StatusNotFound).JSON(helpers.NotFound(err.Error()))
		default:
			return f.Status(fiber.StatusInternalServerError).JSON(helpers.InternalServerError(err.Error()))
		}
	}

	return f.Status(fiber.StatusOK).JSON(helpers.SuccessOK("success delete course", nil))
}
