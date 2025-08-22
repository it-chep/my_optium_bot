package dal

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *Dal {
	return &Dal{
		pool: pool,
	}
}

func (d *Dal) DeleteAdminMessage(ctx context.Context, adminMessageID int64) error {
	sql := `
		delete from admin_messages where id = $1;
	`

	_, err := d.pool.Exec(ctx, sql, adminMessageID)
	return err
}

func (d *Dal) DeleteDoctorMessage(ctx context.Context, doctorMessageID int64) error {
	sql := `
		delete from doctor_messages where id = $1;
	`

	_, err := d.pool.Exec(ctx, sql, doctorMessageID)
	return err
}
