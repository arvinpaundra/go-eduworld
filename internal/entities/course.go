package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type Course struct {
	bun.BaseModel `bun:"table:courses"`

	ID          string `bun:",pk"`
	CategoryId  string `bun:",notnull"`
	UserId      string `bun:",notnull"`
	InterestId  string `bun:",notnull"`
	Title       string `bun:",notnull"`
	Type        string `bun:",notnull"`
	Level       string `bun:",notnull"`
	Description *string
	Thumbnail   *string
	Price       *int
	IsPublished bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `bun:",soft_delete"`
	Category    *Category  `bun:"rel:belongs-to,join:category_id=id"`
	User        *User      `bun:"rel:has-one,join:user_id=id"`
	Interest    *Interest  `bun:"rel:has-one,join:interest_id=id"`
	Modules     []*Module  `bun:"rel:has-many,join:id=course_id"`
}

var _ bun.BeforeAppendModelHook = (*Course)(nil)

func (c *Course) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		c.CreatedAt = time.Now()
		c.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		c.UpdatedAt = time.Now()
	}
	return nil
}
