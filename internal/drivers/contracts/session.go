package contracts

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/uptrace/bun"
)

type SessionRepository interface {
	Begin(ctx context.Context) (*bun.Tx, error)
	FindOne(ctx context.Context, conditions string, args ...any) (*entities.Session, error)
	Save(ctx context.Context, tx *bun.Tx, session *entities.Session) error
	Update(ctx context.Context, tx *bun.Tx, session *entities.Session, conditions string, args ...any) error
	Remove(ctx context.Context, tx *bun.Tx, conditions string, args ...any) error
}
