package session

import (
	"context"
	"database/sql"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type sessionRepository struct {
	conn *bun.DB
}

func NewSQLRepository(conn *bun.DB) contracts.SessionRepository {
	return &sessionRepository{
		conn: conn,
	}
}

func (s *sessionRepository) Begin(ctx context.Context) (*bun.Tx, error) {
	tx, err := s.conn.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return &tx, nil
}

func (s *sessionRepository) FindOne(ctx context.Context, conditions string, args ...any) (*entities.Session, error) {
	var session entities.Session

	err := s.conn.NewSelect().Model(&session).Where(conditions, args...).Scan(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return &session, nil
}

func (s *sessionRepository) Save(ctx context.Context, tx *bun.Tx, session *entities.Session) error {
	_, err := s.conn.NewInsert().Model(session).Conn(tx).Exec(ctx)

	if err != nil {
		return err
	}

	return nil
}

func (s *sessionRepository) Update(ctx context.Context, tx *bun.Tx, session *entities.Session, conditions string, args ...any) error {
	panic("not implemented")
}

func (s *sessionRepository) Remove(ctx context.Context, tx *bun.Tx, conditions string, args ...any) error {
	_, err := s.conn.NewDelete().Model((*entities.Session)(nil)).Conn(tx).Where(conditions, args...).Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}
