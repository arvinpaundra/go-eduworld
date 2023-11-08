package postgres

import (
	"context"
	"fmt"
	"log"

	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
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

func (p *Postgres) Start(ctx context.Context) *bun.DB {
	uri := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.Database,
		p.SSlMode,
	)

	config, err := pgx.ParseConfig(uri)
	if err != nil {
		utils.Logger().Fatal(err)
		return nil
	}

	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook())

	log.Println("connected to postgres")

	return db
}

func Shutdown(ctx context.Context, db *bun.DB) error {
	if err := db.DB.Close(); err != nil {
		utils.Logger().Fatal(err)
	}

	log.Println("success close postgres connection")

	return nil
}
