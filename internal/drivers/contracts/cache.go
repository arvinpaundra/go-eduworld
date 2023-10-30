package contracts

import (
	"context"
	"time"
)

type CacheRepository interface {
	Get(ctx context.Context, key string) (*string, error)
	Set(ctx context.Context, key string, value any, duration time.Duration) error
	Del(ctx context.Context, key string) error
}
