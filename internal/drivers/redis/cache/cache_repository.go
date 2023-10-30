package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/redis/go-redis/v9"
)

type cacheRepository struct {
	rdb *redis.Client
}

func NewKeyValueRepository(rdb *redis.Client) contracts.CacheRepository {
	return &cacheRepository{rdb: rdb}
}

func (c *cacheRepository) Get(ctx context.Context, key string) (*string, error) {
	result, err := c.rdb.Get(ctx, key).Result()

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	if err == redis.Nil {
		utils.Logger().Error(constants.ErrKeyNotFound)
		return nil, constants.ErrKeyNotFound
	}

	return &result, nil
}

func (c *cacheRepository) Set(ctx context.Context, key string, value any, duration time.Duration) error {
	marshaledValue, err := json.Marshal(value)
	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	if err := c.rdb.Set(ctx, key, marshaledValue, duration).Err(); err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (c *cacheRepository) Del(ctx context.Context, key string) error {
	err := c.rdb.Del(ctx, key).Err()

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}
