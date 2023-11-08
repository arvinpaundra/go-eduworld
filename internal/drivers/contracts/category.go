package contracts

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/uptrace/bun"
)

type CategoryRepository interface {
	Find(ctx context.Context) ([]*entities.Category, error)
	FindOne(ctx context.Context, fields string, conditions string, args ...any) (*entities.Category, error)
	Save(ctx context.Context, tx *bun.Tx, course *entities.Category) error
	Update(ctx context.Context, tx *bun.Tx, course *entities.Category, conditions string, args ...any) error
	Remove(ctx context.Context, tx *bun.Tx, conditions string, args ...any) error
}
