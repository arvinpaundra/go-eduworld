package entities

import (
	"context"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
)

type Module struct {
	bun.BaseModel `bun:"table:modules"`

	ID          string `bun:",pk"`
	CourseId    string `bun:",notnull"`
	Title       string `bun:",notnull"`
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time  `bun:",soft_delete"`
	Course      *Course     `bun:"rel:has-one,join:course_id=id"`
	Materials   []*Material `bun:"rel:has-many,join:id=module_id"`
}

var _ bun.BeforeAppendModelHook = (*Module)(nil)

func (m *Module) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now()
		m.UpdatedAt = time.Now()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now()
	}
	return nil
}
