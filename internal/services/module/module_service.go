package module

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/adapters/request"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
)

type ModuleService interface {
	Save(ctx context.Context, input *request.Module, courseId string) error
	Update(ctx context.Context, input *request.Module, courseId string, moduleId string) error
	Remove(ctx context.Context, courseId string, moduleId string) error
}

type moduleService struct {
	courseRepository contracts.CourseRepository
	moduleRepository contracts.ModuleRepository
}

func NewModuleService(
	courseRepository contracts.CourseRepository,
	moduleRepository contracts.ModuleRepository,
) ModuleService {
	return &moduleService{
		courseRepository: courseRepository,
		moduleRepository: moduleRepository,
	}
}

func (m *moduleService) Save(ctx context.Context, input *request.Module, courseId string) error {
	tx, err := m.courseRepository.Begin(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	if _, err := m.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId}); err != nil {
		return err
	}

	newModule := entities.Module{
		ID:          utils.GetID(),
		CourseId:    courseId,
		Title:       input.Title,
		Description: input.Description,
	}

	if err := m.moduleRepository.Save(ctx, tx, &newModule); err != nil {
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

func (m *moduleService) Update(ctx context.Context, input *request.Module, courseId string, moduleId string) error {
	tx, err := m.courseRepository.Begin(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	if _, err := m.moduleRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: moduleId}); err != nil {
		return err
	}

	if _, err := m.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId}); err != nil {
		return err
	}

	updatedModule := entities.Module{
		CourseId:    courseId,
		Title:       input.Title,
		Description: input.Description,
	}

	if err := m.moduleRepository.Update(ctx, tx, &updatedModule, moduleId); err != nil {
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

func (m *moduleService) Remove(ctx context.Context, courseId string, moduleId string) error {
	tx, err := m.courseRepository.Begin(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	if _, err := m.moduleRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: moduleId}); err != nil {
		return err
	}

	if _, err := m.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId}); err != nil {
		return err
	}

	if err := m.moduleRepository.Remove(ctx, tx, moduleId); err != nil {
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
