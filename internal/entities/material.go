package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type Material struct {
	bun.BaseModel `bun:"table:materials"`

	ID          string `bun:",pk"`
	CourseId    string `bun:",notnull"`
	ModuleId    string `bun:",notnull"`
	Title       string `bun:",notnull"`
	Url         string `bun:",notnull"`
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time `bun:",soft_delete"`
}

var _ bun.BeforeAppendModelHook = (*Material)(nil)

// BeforeAppendModel implements schema.BeforeAppendModelHook.
func (m *Material) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
