package course

import (
	"context"
	"database/sql"
	"strings"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/constants"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type courseRepository struct {
	conn *bun.DB
}

func NewSQLRepository(conn *bun.DB) contracts.CourseRepository {
	return &courseRepository{conn: conn}
}

func (c *courseRepository) Begin(ctx context.Context) (*bun.Tx, error) {
	tx, err := c.conn.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return &tx, nil
}

func (c *courseRepository) Find(ctx context.Context, offset int, conditions ...utils.SQLCondition) ([]*entities.Course, error) {
	var courses []*entities.Course

	op := c.conn.NewSelect().
		Model(&courses).
		Relation("Category").
		Relation("Interest").
		Relation("User")

	if len(conditions) > 0 {
		for _, cond := range conditions {
			switch strings.ToUpper(cond.Operator) {
			case "":
				utils.Logger().Error(constants.ErrInvalidSQLOperator)
				return nil, constants.ErrInvalidSQLOperator
			case "LIKE":
				op.Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
					switch cond.Value.(type) {
					case string:
						val := cond.Value.(string)
						if val != "" {
							return op.Where(cond.Column+" LIKE ?", "%"+val+"%")
						}
					default:
						utils.Logger().Error("got non string value")
					}
					return q
				})
			case "=":
				switch cond.Value.(type) {
				case string:
					val := cond.Value.(string)
					op.Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
						if val != "" {
							return q.Where(cond.Column+" = ?", val)
						}

						return q
					})
				default:
					op.Where(cond.Column+" = ?", cond.Value)
				}
			default:
				op.Where(cond.Column+cond.Operator+" ?", cond.Value)
			}
		}
	}

	op.OrderExpr("course.created_at DESC").
		Offset(offset).Limit(15)

	err := op.Scan(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return courses, nil
}

func (c *courseRepository) FindOne(ctx context.Context, fields string, conditions ...utils.SQLCondition) (*entities.Course, error) {
	var course entities.Course

	op := c.conn.NewSelect().
		Model(&course).
		Column(utils.GetSelectedFields(fields)...)

	if len(conditions) > 0 {
		for _, cond := range conditions {
			switch cond.Operator {
			case "":
				utils.Logger().Error(constants.ErrInvalidSQLOperator)
				return nil, constants.ErrInvalidSQLOperator
			default:
				op.Where(cond.Column+" "+cond.Operator+" ?", cond.Value)
			}
		}
	}

	err := op.Scan(ctx)

	if err != nil {
		if err.Error() == constants.ErrBunNotNotFound.Error() {
			utils.Logger().Error(err)
			return nil, constants.ErrCourseNotFound
		}

		utils.Logger().Error(err)
		return nil, err
	}

	return &course, nil
}

func (c *courseRepository) DetailCourse(ctx context.Context, fields string, conditions ...utils.SQLCondition) (*entities.Course, error) {
	var course entities.Course

	op := c.conn.NewSelect().
		Model(&course).
		Relation("Category").
		Relation("Interest").
		Relation("User").
		Relation("Modules", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.OrderExpr("module.created_at ASC")
		}).
		Relation("Modules.Materials", func(q *bun.SelectQuery) *bun.SelectQuery {
			return q.OrderExpr("material.created_at ASC")
		}).
		Column(utils.GetSelectedFields(fields)...)

	if len(conditions) > 0 {
		for _, cond := range conditions {
			switch cond.Operator {
			case "":
				utils.Logger().Error(constants.ErrInvalidSQLOperator)
				return nil, constants.ErrInvalidSQLOperator
			default:
				op.Where(cond.Column+" "+cond.Operator+" ?", cond.Value)
			}
		}
	}

	err := op.Scan(ctx)

	if err != nil {
		if err.Error() == constants.ErrBunNotNotFound.Error() {
			utils.Logger().Error(err)
			return nil, constants.ErrCourseNotFound
		}

		utils.Logger().Error(err)
		return nil, err
	}

	return &course, nil
}

func (c *courseRepository) Save(ctx context.Context, tx *bun.Tx, course *entities.Course) error {
	_, err := c.conn.NewInsert().Model(course).Conn(tx).Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (c *courseRepository) Update(ctx context.Context, tx *bun.Tx, course *entities.Course, courseId string) error {
	_, err := c.conn.NewUpdate().
		Model(course).
		Conn(tx).
		OmitZero().
		Where("id = ?", courseId).
		Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (c *courseRepository) Remove(ctx context.Context, tx *bun.Tx, courseId string) error {
	_, err := c.conn.NewDelete().
		Model((*entities.Course)(nil)).
		Conn(tx).
		Where("id = ?", courseId).
		Exec(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (c *courseRepository) Count(ctx context.Context, conditions ...utils.SQLCondition) (int, error) {
	var total int

	op := c.conn.NewSelect().
		Model((*entities.Course)(nil)).
		ColumnExpr("COUNT(?)", bun.Ident("id"))

	if len(conditions) > 0 {
		for _, cond := range conditions {
			switch strings.ToUpper(cond.Operator) {
			case "":
				utils.Logger().Error(constants.ErrInvalidSQLOperator)
				return 0, constants.ErrInvalidSQLOperator
			case "LIKE":
				op.Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
					switch cond.Value.(type) {
					case string:
						val := cond.Value.(string)
						if val != "" {
							return op.Where(cond.Column+" LIKE ?", "%"+val+"%")
						}
					default:
						utils.Logger().Error("got non string value")
					}
					return q
				})
			case "=":
				switch cond.Value.(type) {
				case string:
					val := cond.Value.(string)
					op.Apply(func(q *bun.SelectQuery) *bun.SelectQuery {
						if val != "" {
							return q.Where(cond.Column+" = ?", val)
						}

						return q
					})
				default:
					op.Where(cond.Column+" = ?", cond.Value)
				}
			default:
				op.Where(cond.Column+" "+cond.Operator+" ?", cond.Value)
			}
		}
	}

	err := op.Scan(ctx, &total)

	if err != nil {
		utils.Logger().Error(err)
		return 0, err
	}

	return total, nil
}
