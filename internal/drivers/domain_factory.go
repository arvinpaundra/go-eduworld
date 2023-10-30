package drivers

import (
	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/device"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/interest"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/session"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/user"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/redis/cache"
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

func NewDeviceRepository(conn *bun.DB) contracts.DeviceRepository {
	return device.NewSQLRepository(conn)
}

func NewSessionRepository(conn *bun.DB) contracts.SessionRepository {
	return session.NewSQLRepository(conn)
}

func NewUserRepository(conn *bun.DB) contracts.UserRepository {
	return user.NewSQLRepository(conn)
}

func NewInterestRepository(conn *bun.DB) contracts.InterestRepository {
	return interest.NewSQLRepository(conn)
}

func NewCacheRepository(rdb *redis.Client) contracts.CacheRepository {
	return cache.NewKeyValueRepository(rdb)
}
