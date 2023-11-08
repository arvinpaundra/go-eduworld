package module

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type moduleRepository struct {
	conn *bun.DB
}

func NewSQLRepository(conn *bun.DB) contracts.ModuleRepository {
	return &moduleRepository{conn: conn}
}

func (m *moduleRepository) FindOne(ctx context.Context, fields string, conditions ...utils.SQLCondition) (*entities.Module, error) {
	var module entities.Module

	op := m.conn.NewSelect().
		Model(&module).
		Column(utils.GetSelectedFields(fields)...)

	if len(conditions) > 0 {
		for _, cond := range conditions {
			switch cond.Operator {
			case "":
				utils.Logger().Error(constants.ErrInvalidSQLOperator)
				return nil, constants.ErrInvalidSQLOperator
			default:
				op.Where(cond.Column+" "+cond.Operator+" ?", cond.Value)
			}
		}
	}

	err := op.Scan(ctx)

	if err != nil {
		if err.Error() == constants.ErrBunNotNotFound.Error() {
			utils.Logger().Error(err)
			return nil, constants.ErrModuleNotFounnd
		}

		utils.Logger().Error(err)
		return nil, err
	}

	return &module, nil
}

func (m *moduleRepository) Save(ctx context.Context, tx *bun.Tx, module *entities.Module) error {
	_, err := m.conn.NewInsert().Model(module).Conn(tx).Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (m *moduleRepository) Update(ctx context.Context, tx *bun.Tx, module *entities.Module, moduleId string) error {
	_, err := m.conn.NewUpdate().
		Model(module).
		Conn(tx).
		OmitZero().
		Where("id = ?", moduleId).
		Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (m *moduleRepository) Remove(ctx context.Context, tx *bun.Tx, moduleId string) error {
	_, err := m.conn.NewDelete().
		Model((*entities.Module)(nil)).
		Conn(tx).
		Where("id = ?", moduleId).
		Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}
