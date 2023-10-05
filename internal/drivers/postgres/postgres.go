package postgres

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
	SSlMode  string
	Timezone string
}

func NewPostgres(postgres *Postgres) *Postgres {
	return postgres
}

func (p *Postgres) Start(ctx context.Context) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		p.Host,
		p.User,
		p.Password,
		p.Database,
		p.Port,
		p.SSlMode,
		p.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Fatalf("error while connect to postgres: %e", err)
	}

	log.Println("connected to postgres")

	return db
}

func Shutdown(ctx context.Context, db *gorm.DB) error {
	postgres, err := db.DB()

	if err != nil {
		log.Fatalf("error while getting postgres instance: %e", err)
	}

	if err := postgres.Close(); err != nil {
		log.Fatalf("error while closing postgres connection: %e", err)
	}

	log.Println("postgres connection closed")

	return nil
}
