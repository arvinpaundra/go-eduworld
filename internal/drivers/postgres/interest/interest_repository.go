package interest

import (
	"context"
	"database/sql"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type interestRepository struct {
	conn *bun.DB
}

func NewSQLRepository(conn *bun.DB) contracts.InterestRepository {
	return &interestRepository{
		conn: conn,
	}
}

func (i *interestRepository) Begin(ctx context.Context) (*bun.Tx, error) {
	tx, err := i.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (i *interestRepository) Find(ctx context.Context, keyword string, limit int, offset int) ([]*entities.Interest, error) {
	var interests []*entities.Interest

	_, err := i.conn.NewSelect().Model(&interests).ExcludeColumn("created_at", "updated_at").
		Where("name LIKE ?", "%"+keyword+"%").Limit(limit).Offset(offset).Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return interests, nil
}

func (i *interestRepository) FindOne(ctx context.Context, fields string, conditions string, args ...any) (*entities.Interest, error) {
	var interest entities.Interest

	err := i.conn.NewSelect().Model(&interest).Column(utils.GetSelectedFields(fields)...).Where(conditions, args...).Scan(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return &interest, nil
}

func (i *interestRepository) Save(ctx context.Context, tx *bun.Tx, interest *entities.Interest) error {
	panic("not implemented")
}

func (i *interestRepository) Update(ctx context.Context, tx *bun.Tx, interest *entities.Interest, conditions string, args ...any) error {
	panic("not implemented")
}

func (i *interestRepository) Count(ctx context.Context, keyword string) (int, error) {
	panic("not implemented")
}
