package material

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/adapters/request"
	"github.com/arvinpaundra/go-eduworld/internal/adapters/response"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
)

type MaterialService interface {
	FindOne(ctx context.Context, courseId string, moduleId string, materialId string) (*response.Material, error)
	Save(ctx context.Context, input *request.Material, courseId string, moduleId string) error
	Update(ctx context.Context, input *request.Material, courseId string, moduleId string, materialId string) error
	Remove(ctx context.Context, courseId string, moduleId string, materialId string) error
}

type materialService struct {
	courseRepository   contracts.CourseRepository
	moduleRepository   contracts.ModuleRepository
	materialRepository contracts.MaterialRepository
}

func NewMaterialService(
	courseRepository contracts.CourseRepository,
	moduleRepository contracts.ModuleRepository,
	materialRepository contracts.MaterialRepository,
) MaterialService {
	return &materialService{
		courseRepository:   courseRepository,
		moduleRepository:   moduleRepository,
		materialRepository: materialRepository,
	}
}

func (m *materialService) FindOne(ctx context.Context, courseId string, moduleId string, materialId string) (*response.Material, error) {
	if _, err := m.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId}); err != nil {
		return nil, err
	}

	if _, err := m.moduleRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: moduleId}); err != nil {
		return nil, err
	}

	material, err := m.materialRepository.FindOne(ctx, "*", utils.SQLCondition{Column: "id", Operator: "=", Value: materialId})

	if err != nil {
		return nil, err
	}

	result := response.Material{
		ID:          material.ID,
		CourseId:    material.CourseId,
		ModuleId:    material.ModuleId,
		Title:       material.Title,
		Url:         material.Url,
		Description: material.Description,
	}

	return &result, nil
}

func (m *materialService) Save(ctx context.Context, input *request.Material, courseId string, moduleId string) error {
	tx, err := m.courseRepository.Begin(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	if _, err := m.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId}); err != nil {
		return err
	}

	if _, err := m.moduleRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: moduleId}); err != nil {
		return err
	}

	newMaterial := entities.Material{
		ID:          utils.GetID(),
		CourseId:    courseId,
		ModuleId:    moduleId,
		Title:       input.Title,
		Url:         "",
		Description: input.Description,
	}

	if err := m.materialRepository.Save(ctx, tx, &newMaterial); err != nil {
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

func (m *materialService) Update(ctx context.Context, input *request.Material, courseId string, moduleId string, materialId string) error {
	tx, err := m.courseRepository.Begin(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	if _, err := m.materialRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: materialId}); err != nil {
		return err
	}

	if _, err := m.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId}); err != nil {
		return err
	}

	if _, err := m.moduleRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: moduleId}); err != nil {
		return err
	}

	updatedMaterial := entities.Material{
		CourseId:    courseId,
		ModuleId:    moduleId,
		Title:       input.Title,
		Url:         "",
		Description: input.Description,
	}

	if err := m.materialRepository.Update(ctx, tx, &updatedMaterial, materialId); err != nil {
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

func (m *materialService) Remove(ctx context.Context, courseId string, moduleId string, materialId string) error {
	tx, err := m.courseRepository.Begin(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	if _, err := m.materialRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: materialId}); err != nil {
		return err
	}

	if _, err := m.courseRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: courseId}); err != nil {
		return err
	}

	if _, err := m.moduleRepository.FindOne(ctx, "id", utils.SQLCondition{Column: "id", Operator: "=", Value: moduleId}); err != nil {
		return err
	}

	if err := m.materialRepository.Remove(ctx, tx, materialId); err != nil {
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
