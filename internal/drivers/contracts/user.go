package contracts

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	Begin(ctx context.Context) (*bun.Tx, error)
	FindOne(ctx context.Context, fields string, conditions string, args ...any) (*entities.User, error)
	Save(ctx context.Context, tx *bun.Tx, user *entities.User) error
	Update(ctx context.Context, tx *bun.Tx, user *entities.User, conditions string, args ...any) error
}
