package course

import (
	"context"
	"strings"
	"time"

	"github.com/arvinpaundra/go-eduworld/internal/adapters/request"
	"github.com/arvinpaundra/go-eduworld/internal/adapters/response"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/helpers"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
)

type CourseService interface {
	Find(ctx context.Context, query *request.QueryFindCourses) ([]*response.Course, *helpers.OffsetPagination, error)
	FindByMentor(ctx context.Context, mentorId string, query *request.QueryFindCourses) ([]*response.Course, *helpers.OffsetPagination, error)
	DetailCourse(ctx context.Context, courseId string) (*response.Course, error)
	Save(ctx context.Context, input *request.Course) error
	Update(ctx context.Context, input *request.Course, courseId string) error
	Remove(ctx context.Context, courseId string) error
}

type courseService struct {
	userRepository     contracts.UserRepository
	interestRepository contracts.InterestRepository
	categoryRepository contracts.CategoryRepository
	courseRepository   contracts.CourseRepository
	moduleRepository   contracts.ModuleRepository
	materialRepository contracts.MaterialRepository
	cacheRepository    contracts.CacheRepository
}

func NewCourseService(
	userRepository contracts.UserRepository,
	interestRepository contracts.InterestRepository,
	categoryRepository contracts.CategoryRepository,
	courseRepository contracts.CourseRepository,
	moduleRepository contracts.ModuleRepository,
	materialRepository contracts.MaterialRepository,
	cacheRepository contracts.CacheRepository,
) CourseService {
	return &courseService{
		userRepository:     userRepository,
		interestRepository: interestRepository,
		categoryRepository: categoryRepository,
		courseRepository:   courseRepository,
		moduleRepository:   moduleRepository,
		materialRepository: materialRepository,
		cacheRepository:    cacheRepository,
	}
}

func (c *courseService) Find(ctx context.Context, query *request.QueryFindCourses) ([]*response.Course, *helpers.OffsetPagination, error) {
	courses, err := c.courseRepository.Find(
		ctx,
		query.Page,
		utils.SQLCondition{
			Column:   "LOWER(course.title)",
			Operator: "like",
			Value:    strings.ToLower(query.Keyword),
		},
		utils.SQLCondition{
			Column:   "course.is_published",
			Operator: "=",
			Value:    true,
		},
		utils.SQLCondition{
			Column:   "course.category_id",
			Operator: "=",
			Value:    query.Category,
		},
	)

	if err != nil {
		return nil, nil, err
	}

	count, err := c.courseRepository.Count(
		ctx,
		utils.SQLCondition{
			Column:   "LOWER(course.title)",
			Operator: "like",
			Value:    strings.ToLower(query.Keyword),
		},
		utils.SQLCondition{
			Column:   "course.category_id",
			Operator: "=",
			Value:    query.Category,
		},
		utils.SQLCondition{
			Column:   "course.is_published",
			Operator: "=",
			Value:    true,
		},
	)

	if err != nil {
		return nil, nil, err
	}

	return response.ToResponseCourses(courses), helpers.NewOffsetPagination(query.Page, 15, count), nil
}

func (c *courseService) FindByMentor(ctx context.Context, mentorId string, query *request.QueryFindCourses) ([]*response.Course, *helpers.OffsetPagination, error) {
	courses, err := c.courseRepository.Find(
		ctx,
		query.Page,
		utils.SQLCondition{
			Column:   "LOWER(course.title)",
			Operator: "like",
			Value:    strings.ToLower(query.Keyword),
		},
		utils.SQLCondition{
			Column:   "course.user_id",
			Operator: "=",
			Value:    mentorId,
		},
		utils.SQLCondition{
			Column:   "course.category_id",
			Operator: "=",
			Value:    query.Category,
		},
	)

	if err != nil {
		return nil, nil, err
	}

	count, err := c.courseRepository.Count(
		ctx,
		utils.SQLCondition{
			Column:   "LOWER(course.title)",
			Operator: "like",
			Value:    strings.ToLower(query.Keyword),
		},
		utils.SQLCondition{
			Column:   "course.user_id",
			Operator: "=",
			Value:    mentorId,
		},
		utils.SQLCondition{
			Column:   "course.category_id",
			Operator: "=",
			Value:    query.Category,
		},
	)

	if err != nil {
		return nil, nil, err
	}

	return response.ToResponseCourses(courses), helpers.NewOffsetPagination(query.Page, 15, count), nil
}

func (c *courseService) DetailCourse(ctx context.Context, courseId string) (*response.Course, error) {
	course, err := c.courseRepository.DetailCourse(
		ctx,
		"course.id,course.category_id,course.user_id,course.interest_id,course.title,course.is_published,course.thumbnail,course.created_at,course.updated_at",
		utils.SQLCondition{
			Column:   "course.id",
			Operator: "=",
			Value:    courseId,
		},
	)

	if err != nil {
		return nil, err
	}

	return response.ToResponseCourse(course), nil
}

func (c *courseService) Save(ctx context.Context, input *request.Course) error {
	tx, err := c.courseRepository.Begin(ctx)

	if err != nil {
		return err
	}

	if _, err := c.categoryRepository.FindOne(ctx, "id", "id=?", input.CategoryId); err != nil {
		return err
	}

	if _, err := c.interestRepository.FindOne(ctx, "id", "id=?", input.InterestId); err != nil {
		return err
	}

	// save into database
	newCourse := entities.Course{
		ID:          utils.GetID(),
		CategoryId:  input.CategoryId,
		UserId:      input.UserId,
		InterestId:  input.InterestId,
		Title:       input.Title,
		Type:        input.Type,
		Level:       input.Level,
		Price:       input.Price,
		Description: input.Description,
		Thumbnail:   nil,
		IsPublished: input.IsPublished,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := c.courseRepository.Save(ctx, tx, &newCourse); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			utils.Logger().Error(errorRollback)
			return errorRollback
		}

		return err
	}

	if errorCommit := tx.Commit(); errorCommit != nil {
		utils.Logger().Error(errorCommit)
		return errorCommit
	}

	return nil
}

func (c *courseService) Update(ctx context.Context, input *request.Course, courseId string) error {
	tx, err := c.courseRepository.Begin(ctx)

	if err != nil {
		return err
	}

	if _, err := c.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId}); err != nil {
		return err
	}

	if _, err := c.categoryRepository.FindOne(ctx, "id", "id=?", input.CategoryId); err != nil {
		return err
	}

	if _, err := c.userRepository.FindOne(ctx, "id", "id=?", input.UserId); err != nil {
		return err
	}

	if _, err := c.interestRepository.FindOne(ctx, "id", "id=?", input.InterestId); err != nil {
		return err
	}

	// update into database
	updatedCourse := entities.Course{
		ID:          utils.GetID(),
		CategoryId:  input.CategoryId,
		UserId:      input.UserId,
		InterestId:  input.InterestId,
		Title:       input.Title,
		Type:        input.Type,
		Level:       input.Level,
		Description: input.Description,
		Price:       input.Price,
		Thumbnail:   nil,
		IsPublished: input.IsPublished,
		UpdatedAt:   time.Now(),
	}

	if err := c.courseRepository.Update(ctx, tx, &updatedCourse, courseId); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			utils.Logger().Error(errorRollback)
			return errorRollback
		}

		return err
	}

	if errorCommit := tx.Commit(); errorCommit != nil {
		utils.Logger().Error(errorCommit)
		return errorCommit
	}

	return nil
}

func (c *courseService) Remove(ctx context.Context, courseId string) error {
	tx, err := c.courseRepository.Begin(ctx)

	if err != nil {
		return err
	}

	course, err := c.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId})
	if err != nil {
		return err
	}

	if err := c.courseRepository.Remove(ctx, tx, course.ID); err != nil {
		if errorRollback := tx.Rollback(); errorRollback != nil {
			utils.Logger().Error(errorRollback)
			return errorRollback
		}

		return err
	}

	if errorCommit := tx.Commit(); errorCommit != nil {
		utils.Logger().Error(errorCommit)
		return errorCommit
	}

	return nil
}
