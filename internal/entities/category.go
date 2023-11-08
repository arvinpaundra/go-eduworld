package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type Category struct {
	bun.BaseModel `bun:"table:categories"`

	ID          string `bun:",pk"`
	Name        string `bun:",notnull"`
	Description *string
	Image       *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `bun:",soft_delete"`
}

var _ bun.BeforeAppendModelHook = (*Category)(nil)

func (c *Category) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		c.CreatedAt = time.Now()
		c.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		c.UpdatedAt = time.Now()
	}
	return nil
}
