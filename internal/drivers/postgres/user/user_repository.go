package user

import (
	"context"
	"database/sql"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type userRepository struct {
	conn *bun.DB
}

func NewSQLRepository(conn *bun.DB) contracts.UserRepository {
	return &userRepository{
		conn: conn,
	}
}

func (u *userRepository) Begin(ctx context.Context) (*bun.Tx, error) {
	tx, err := u.conn.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return &tx, nil
}

func (u *userRepository) FindOne(ctx context.Context, fields string, conditions string, args ...any) (*entities.User, error) {
	var user entities.User

	err := u.conn.NewSelect().Model(&user).Column(utils.GetSelectedFields(fields)...).Where(conditions, args...).Scan(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return &user, nil
}

func (u *userRepository) Save(ctx context.Context, tx *bun.Tx, user *entities.User) error {
	if _, err := u.conn.NewInsert().Model(user).Conn(tx).Exec(ctx); err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (u *userRepository) Update(ctx context.Context, tx *bun.Tx, user *entities.User, conditions string, args ...any) error {
	if _, err := u.conn.NewUpdate().Model(user).Conn(tx).OmitZero().Where(conditions, args...).Exec(ctx); err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}
