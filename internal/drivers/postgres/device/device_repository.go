package device

import (
	"context"
	"database/sql"

	"github.com/arvinpaundra/go-eduworld/internal/drivers/contracts"
	"github.com/arvinpaundra/go-eduworld/internal/entities"
	"github.com/arvinpaundra/go-eduworld/pkg/utils"
	"github.com/uptrace/bun"
)

type deviceRepository struct {
	conn *bun.DB
}

func NewSQLRepository(conn *bun.DB) contracts.DeviceRepository {
	return &deviceRepository{
		conn: conn,
	}
}

func (d *deviceRepository) Begin(ctx context.Context) (*bun.Tx, error) {
	tx, err := d.conn.BeginTx(ctx, &sql.TxOptions{})

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return &tx, nil
}

func (d *deviceRepository) FindOne(ctx context.Context, fields string, conditions string, args ...any) (*entities.Device, error) {
	var device entities.Device
	cols := utils.GetSelectedFields(fields)

	err := d.conn.NewSelect().Model(&device).Column(cols...).Where(conditions, args...).Scan(ctx)

	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return &device, nil
}

func (d *deviceRepository) Save(ctx context.Context, tx *bun.Tx, device *entities.Device) error {
	if _, err := d.conn.NewInsert().Model(device).Conn(tx).Exec(ctx); err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (d *deviceRepository) Update(ctx context.Context, tx *bun.Tx, device *entities.Device, conditions string, args ...any) error {
	_, err := d.conn.NewUpdate().Model(device).Conn(tx).OmitZero().Where(conditions, args...).Exec(ctx)
	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (d *deviceRepository) Remove(ctx context.Context, tx *bun.Tx, conditions string, args ...any) error {
	_, err := d.conn.NewDelete().Model((*entities.Device)(nil)).Conn(tx).Where(conditions, args...).Exec(ctx)
	if err != nil {
		utils.Logger().Error(err)
		return err
	}

	return nil
}

func (d *deviceRepository) Count(ctx context.Context, conditions string, args ...any) (int, error) {
	var total int
	_, err := d.conn.NewSelect().Model((*entities.Device)(nil)).
		ColumnExpr("COUNT(?)", bun.Ident("id")).Where(conditions, args...).Exec(ctx, &total)

	if err != nil {
		utils.Logger().Error(err)
		return 0, err
	}

	return total, nil
}
