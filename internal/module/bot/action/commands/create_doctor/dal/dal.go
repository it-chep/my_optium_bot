package create_doctor_dal

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
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

func (d *Dal) UpsertDoctor(ctx context.Context, doctor bot_dto.User) error {
	sql := `insert into doctors (tg_id, full_name, tg_username) 
				values ($1, $2, $3)
			on conflict (tg_id) do update 
			set 
			    full_name = excluded.full_name,
				tg_username = excluded.tg_username
			`
	args := []interface{}{
		doctor.ID,
		doctor.Name,
		doctor.UserName,
	}

	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}
