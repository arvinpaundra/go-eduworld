package material

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type materialRepository struct {
	conn *bun.DB
}

func NewSQLRepository(conn *bun.DB) contracts.MaterialRepository {
	return &materialRepository{conn: conn}
}

func (m *materialRepository) FindOne(ctx context.Context, fields string, conditions ...utils.SQLCondition) (*entities.Material, error) {
	var material entities.Material

	op := m.conn.NewSelect().
		Model(&material).
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
			return nil, constants.ErrMaterialNotFound
		}

		utils.Logger().Error(err)
		return nil, err
	}

	return &material, nil
}

func (m *materialRepository) Save(ctx context.Context, tx *bun.Tx, material *entities.Material) error {
	_, err := m.conn.NewInsert().Model(material).Conn(tx).Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (m *materialRepository) Update(ctx context.Context, tx *bun.Tx, material *entities.Material, materialId string) error {
	_, err := m.conn.NewUpdate().
		Model(material).
		Conn(tx).
		OmitZero().
		Where("id = ?", materialId).
		Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (m *materialRepository) Remove(ctx context.Context, tx *bun.Tx, materialId string) error {
	_, err := m.conn.NewDelete().
		Model((*entities.Material)(nil)).
		Conn(tx).
		Where("id = ?", materialId).
		Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}
