package redis

import (
	"context"
	"log"

	// "crypto/tls"
	"fmt"

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
		log.Printf("error while connect to redis: %e", err)
	}

	log.Println("connected to redis")

	return rdb
}

func Shutdown(ctx context.Context, rc *redis.Client) error {
	if err := rc.Close(); err != nil {
		log.Printf("error while closing redis connection: %e", err)
	}

	log.Println("redis connection closed")

	return nil
}
