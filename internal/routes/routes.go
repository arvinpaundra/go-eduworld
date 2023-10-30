package routes

import (
	authCtrl "github.com/arvinpaundra/go-eduworld/internal/controllers/auth"
	authSrvc "github.com/arvinpaundra/go-eduworld/internal/services/auth"
	"github.com/uptrace/bun"

	"github.com/arvinpaundra/go-eduworld/internal/drivers"
	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Sql      *bun.DB
	Redis    *redis.Client
	Fiber    *fiber.App
	RabbitMQ *amqp091.Connection
}

func (c *Config) Start() {
	// dependecy injection
	deviceRepository := drivers.NewDeviceRepository(c.Sql)
	sessionRepository := drivers.NewSessionRepository(c.Sql)
	userRepository := drivers.NewUserRepository(c.Sql)
	interestRepository := drivers.NewInterestRepository(c.Sql)
	cacheRepository := drivers.NewCacheRepository(c.Redis)

	authService := authSrvc.NewAuthService(
		userRepository,
		deviceRepository,
		sessionRepository,
		interestRepository,
		cacheRepository,
	)

	authController := authCtrl.NewAuthController(authService)

	// group app version
	v1 := c.Fiber.Group("/api/v1")

	// auth routes
	auth := v1.Group("/auth")

	auth.Post("/login", authController.HandlerLogin)
	auth.Post("/register", authController.HandlerRegister)
	auth.Post("/logout", authController.HandlerLogout)
}
