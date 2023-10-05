package rabbitmq

import (
	"context"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Host     string
	Username string
	Password string
	Port     string
}

func NewRabbitMQ(rabbitmq *RabbitMQ) *RabbitMQ {
	return rabbitmq
}

func (rmq *RabbitMQ) Start(ctx context.Context) *amqp091.Connection {
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rmq.Username, rmq.Password, rmq.Host, rmq.Port)

	conn, err := amqp091.Dial(connStr)

	if err != nil {
		log.Printf("error while connect to rabbitmq instance: %e", err)
	}

	log.Println("connected to rabbitmq instance")

	return conn
}

func Shutdown(ctx context.Context, conn *amqp091.Connection) error {
	if err := conn.Close(); err != nil {
		log.Printf("error while closing rabbitmq connection: %e", err)
	}

	log.Println("rabbitmq connection closed")

	return nil
}
