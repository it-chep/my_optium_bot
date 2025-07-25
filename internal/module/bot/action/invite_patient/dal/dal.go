package invite_patient_dal

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

func (d *Dal) SetUserTgID(ctx context.Context, tgID, chatID int64) error {
	sql := `
		with chat as (
    		update patient_doctor set patient_tg = $1 where chat_id = $2 and patient_tg is null returning patient_id
		)
		update patients set tg_id = $1 where id = (select patient_id from chat)
	`
	args := []interface{}{tgID, chatID}
	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}
