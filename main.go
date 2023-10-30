package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/arvinpaundra/go-eduworld/internal/configs"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/redis"
	"github.com/arvinpaundra/go-eduworld/internal/routes"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()
	constants.Validate = validator.New()

	// start postgres connection
	pg := postgres.NewPostgres(&postgres.Postgres{
		Host:     configs.GetConfig("POSTGRES_HOST"),
		User:     configs.GetConfig("POSTGRES_USER"),
		Password: configs.GetConfig("POSTGRES_PASSWORD"),
		Database: configs.GetConfig("POSTGRES_NAME"),
		Port:     configs.GetConfig("POSTGRES_PORT"),
		SSlMode:  configs.GetConfig("POSTGRES_SSLMODE"),
		Timezone: configs.GetConfig("POSTGRES_TIMEZONE"),
	}).Start(ctx)

	// start redis connection
	rdb := redis.NewRedis(&redis.Redis{
		Host:     configs.GetConfig("REDIS_HOST"),
		User:     configs.GetConfig("REDIS_USERNAME"),
		Password: configs.GetConfig("REDIS_PASSWORD"),
		Port:     configs.GetConfig("REDIS_PORT"),
	}).Start(ctx)

	// start rabbitmq instance
	// rmq := rabbitmq.NewRabbitMQ(&rabbitmq.RabbitMQ{
	// 	Host:     configs.GetConfig("RABBITMQ_HOST"),
	// 	Username: configs.GetConfig("RABBITMQ_USERNAME"),
	// 	Password: configs.GetConfig("RABBITMQ_PASSWORD"),
	// 	Port:     configs.GetConfig("RABBITMQ_PORT"),
	// }).Start(ctx)

	// create new gin instance
	app := fiber.New()

	// any middlewares goes here
	// app.Use()

	routeConfig := routes.Config{
		Sql:   pg,
		Redis: rdb,
		Fiber: app,
		// RabbitMQ: rmq,
	}

	routeConfig.Start()

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := app.Listen(configs.GetConfig("APP_ADDR")); err != nil && err != http.ErrServerClosed {
			log.Fatalf("shutting down server. err: %e", err)
		}
	}()

	// perform graceful shutdown to entire applications
	wait := utils.GracefulShutdown(ctx, time.Second*5, map[string]utils.Operation{
		"postgres": func(ctx context.Context) error {
			return postgres.Shutdown(ctx, pg)
		},
		"redis": func(ctx context.Context) error {
			return redis.Shutdown(ctx, rdb)
		},
		// "rabbitmq": func(ctx context.Context) error {
		// 	return rabbitmq.Shutdown(ctx, rmq)
		// },
		"http-server": func(ctx context.Context) error {
			return app.ShutdownWithContext(ctx)
		},
	})

	<-wait
}
