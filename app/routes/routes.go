package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Config struct {
	Sql      *gorm.DB
	Redis    *redis.Client
	Gin      *gin.Engine
	RabbitMQ *amqp091.Connection
}

func Start(config *Config) {

}
