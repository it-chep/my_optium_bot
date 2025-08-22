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

func (d *Dal) CreateUserList(ctx context.Context, userID, listID int64) error {
	sql := `
		insert into users_lists (user_id, list_id) values ($1, $2)
		`
	_, err := d.pool.Exec(ctx, sql, userID, listID)
	return err
}
