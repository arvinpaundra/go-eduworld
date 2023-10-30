package contracts

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/uptrace/bun"
)

type DeviceRepository interface {
	Begin(ctx context.Context) (*bun.Tx, error)
	FindOne(ctx context.Context, fields string, conditions string, args ...any) (*entities.Device, error)
	Save(ctx context.Context, tx *bun.Tx, device *entities.Device) error
	Update(ctx context.Context, tx *bun.Tx, device *entities.Device, conditions string, args ...any) error
	Remove(ctx context.Context, tx *bun.Tx, conditions string, args ...any) error
	Count(ctx context.Context, conditions string, args ...any) (int, error)
}
