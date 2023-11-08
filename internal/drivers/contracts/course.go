package contracts

import (
	"context"

	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type CourseRepository interface {
	Begin(ctx context.Context) (*bun.Tx, error)
	Find(ctx context.Context, offset int, conditions ...utils.SQLCondition) ([]*entities.Course, error)
	FindOne(ctx context.Context, fields string, conditions ...utils.SQLCondition) (*entities.Course, error)
	DetailCourse(ctx context.Context, fields string, conditions ...utils.SQLCondition) (*entities.Course, error)
	Save(ctx context.Context, tx *bun.Tx, course *entities.Course) error
	Update(ctx context.Context, tx *bun.Tx, course *entities.Course, courseId string) error
	Remove(ctx context.Context, tx *bun.Tx, courseId string) error
	Count(ctx context.Context, conditions ...utils.SQLCondition) (int, error)
}
