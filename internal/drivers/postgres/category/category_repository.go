package category

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type categoryRepository struct {
	conn *bun.DB
}

func NewSQLRepository(conn *bun.DB) contracts.CategoryRepository {
	return &categoryRepository{conn: conn}
}

func (c *categoryRepository) Find(ctx context.Context) ([]*entities.Category, error) {
	panic("not implemented")
}

func (c *categoryRepository) FindOne(ctx context.Context, fields string, conditions string, args ...any) (*entities.Category, error) {
	var category entities.Category

	err := c.conn.NewSelect().
		Model(&category).
		Column(utils.GetSelectedFields(fields)...).
		Where(conditions, args...).
		Scan(ctx)

	if err != nil {
		if err.Error() == constants.ErrBunNotNotFound.Error() {
			utils.Logger().Error(err)
			return nil, constants.ErrCategoryNotFound
		}

		utils.Logger().Error(err)
		return nil, err
	}

	return &category, nil
}

func (c *categoryRepository) Save(ctx context.Context, tx *bun.Tx, course *entities.Category) error {
	_, err := c.conn.NewInsert().Model(course).Conn(tx).Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (c *categoryRepository) Update(ctx context.Context, tx *bun.Tx, course *entities.Category, conditions string, args ...any) error {
	_, err := c.conn.NewUpdate().Model(course).Conn(tx).Where(conditions, args...).Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (c *categoryRepository) Remove(ctx context.Context, tx *bun.Tx, conditions string, args ...any) error {
	_, err := c.conn.NewDelete().
		Model((*entities.Category)(nil)).
		Conn(tx).Where(conditions, args...).
		Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}
