package contracts

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type ModuleRepository interface {
	FindOne(ctx context.Context, fields string, conditions ...utils.SQLCondition) (*entities.Module, error)
	Save(ctx context.Context, tx *bun.Tx, module *entities.Module) error
	Update(ctx context.Context, tx *bun.Tx, module *entities.Module, moduleId string) error
	Remove(ctx context.Context, tx *bun.Tx, moduleId string) error
}
