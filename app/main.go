package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/arvinpaundra/go-eduworld/app/routes"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/postgres"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/rabbitmq"
	"github.com/arvinpaundra/go-eduworld/internal/drivers/redis"
	"github.com/arvinpaundra/go-eduworld/internal/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()

	// start postgres connection
	pg := postgres.NewPostgres(&postgres.Postgres{
		Host:     helpers.GetConfig("POSTGRES_HOST"),
		User:     helpers.GetConfig("POSTGRES_USER"),
		Password: helpers.GetConfig("POSTGRES_PASSWORD"),
		Database: helpers.GetConfig("POSTGRES_NAME"),
		Port:     helpers.GetConfig("POSTGRES_PORT"),
		SSlMode:  helpers.GetConfig("POSTGRES_SSLMODE"),
		Timezone: helpers.GetConfig("POSTGRES_TIMEZONE"),
	}).Start(ctx)

	// start redis connection
	rdb := redis.NewRedis(&redis.Redis{
		Host:     helpers.GetConfig("REDIS_HOST"),
		User:     helpers.GetConfig("REDIS_USERNAME"),
		Password: helpers.GetConfig("REDIS_PASSWORD"),
		Port:     helpers.GetConfig("REDIS_PORT"),
	}).Start(ctx)

	// start rabbitmq instance
	rmq := rabbitmq.NewRabbitMQ(&rabbitmq.RabbitMQ{
		Host:     helpers.GetConfig("RABBITMQ_HOST"),
		Username: helpers.GetConfig("RABBITMQ_USERNAME"),
		Password: helpers.GetConfig("RABBITMQ_PASSWORD"),
		Port:     helpers.GetConfig("RABBITMQ_PORT"),
	}).Start(ctx)

	// set application mode
	if helpers.GetConfig("APP_MODE") == "development" {
		gin.SetMode(gin.DebugMode)
	}
	if helpers.GetConfig("APP_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	if helpers.GetConfig("APP_MODE") == "test" {
		gin.SetMode(gin.TestMode)
	}

	// create new gin instance
	g := gin.New()

	// any middlewares goes here
	g.ForwardedByClientIP = true
	g.SetTrustedProxies([]string{"127.0.0.1"})
	g.Use()

	routes.Start(&routes.Config{
		Sql:      pg,
		Redis:    rdb,
		Gin:      g,
		RabbitMQ: rmq,
	})

	srv := &http.Server{
		Addr:    helpers.GetConfig("APP_PORT"),
		Handler: g,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("shutting down server")
		}
	}()

	// perform graceful shutdown to entire applications
	wait := helpers.GracefulShutdown(ctx, time.Second*5, map[string]helpers.Operation{
		"postgres": func(ctx context.Context) error {
			return postgres.Shutdown(ctx, pg)
		},
		"redis": func(ctx context.Context) error {
			return redis.Shutdown(ctx, rdb)
		},
		"rabbitmq": func(ctx context.Context) error {
			return rabbitmq.Shutdown(ctx, rmq)
		},
		"http-server": func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	<-wait
}
