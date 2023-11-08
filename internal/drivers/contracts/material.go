package contracts

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type MaterialRepository interface {
	FindOne(ctx context.Context, fields string, conditions ...utils.SQLCondition) (*entities.Material, error)
	Save(ctx context.Context, tx *bun.Tx, material *entities.Material) error
	Update(ctx context.Context, tx *bun.Tx, material *entities.Material, materialId string) error
	Remove(ctx context.Context, tx *bun.Tx, materialId string) error
}
