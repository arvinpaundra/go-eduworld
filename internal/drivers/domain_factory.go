package drivers

import (
	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/category"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/course"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/device"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/interest"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/material"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres/module"
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

func NewCategoryRepository(conn *bun.DB) contracts.CategoryRepository {
	return category.NewSQLRepository(conn)
}

func NewCourseRepository(conn *bun.DB) contracts.CourseRepository {
	return course.NewSQLRepository(conn)
}

func NewModuleRepository(conn *bun.DB) contracts.ModuleRepository {
	return module.NewSQLRepository(conn)
}

func NewMaterialRepository(conn *bun.DB) contracts.MaterialRepository {
	return material.NewSQLRepository(conn)
}
