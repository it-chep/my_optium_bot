package create_admin_chat_dal

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

func (d *Dal) UpsertAdminChat(ctx context.Context, chatID int64) error {
	sql := `insert into admin_chat (chat_id) 
				values ($1)
			on conflict (chat_id) do nothing
			`
	_, err := d.pool.Exec(ctx, sql, chatID)
	return err
}
