package contracts

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/uptrace/bun"
)

type InterestRepository interface {
	Begin(ctx context.Context) (*bun.Tx, error)
	Find(ctx context.Context, keyword string, limit int, offset int) ([]*entities.Interest, error)
	FindOne(ctx context.Context, fields string, conditions string, args ...any) (*entities.Interest, error)
	Save(ctx context.Context, tx *bun.Tx, interest *entities.Interest) error
	Update(ctx context.Context, tx *bun.Tx, interest *entities.Interest, conditions string, args ...any) error
	Count(ctx context.Context, keyword string) (int, error)
}
