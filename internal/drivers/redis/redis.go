package redis

import (
	"context"
	"log"

	// "crypto/tls"
	"fmt"

	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Host     string
	User     string
	Password string
	Port     string
}

func NewRedis(redis *Redis) *Redis {
	return redis
}

func (r *Redis) Start(ctx context.Context) *redis.Client {
	addr := fmt.Sprintf("%s:%s", r.Host, r.Port)

	rdb := redis.NewClient(&redis.Options{
		Username: r.User,
		Addr:     addr,
		Password: r.Password,
		DB:       0,
		// TLSConfig: &tls.Config{},
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		utils.Logger().Fatal(err)
		return nil
	}

	log.Println("connected to redis")

	return rdb
}

func Shutdown(ctx context.Context, rc *redis.Client) error {
	if err := rc.Close(); err != nil {
		utils.Logger().Fatal(err)
	}

	log.Println("success close redis connection")

	return nil
}
