package entities

import (
	"time"

	"github.com/uptrace/bun"
)

type Interest struct {
	bun.BaseModel `bun:"table:interests"`

	ID          string `bun:",pk"`
	Name        string `bun:",notnull"`
	Icon        *string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
