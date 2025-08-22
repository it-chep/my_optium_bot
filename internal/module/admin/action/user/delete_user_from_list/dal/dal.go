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

func (d *Dal) DeleteUserFromList(ctx context.Context, userID, listID int64) error {
	sql := `
	delete from users_lists where user_id = $1 and list_id = $2
	`

	_, err := d.pool.Exec(ctx, sql, userID, listID)
	return err
}
