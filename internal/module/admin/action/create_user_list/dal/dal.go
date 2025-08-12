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

func (d *Dal) CreateUserList(ctx context.Context, listName string) error {
	sql := `
		insert into user_lists (name) values ($1)
		`
	_, err := d.pool.Exec(ctx, sql, listName)
	return err
}
