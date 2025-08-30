package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dal struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Dal {
	return &Dal{pool: pool}
}

func (d *Dal) GetRecipientsCount(ctx context.Context, ids []int64) (int64, error) {
	sql := `
		select count(distinct user_id)
		from users_lists
		where list_id = any ($1::bigint[])
    `
	var count int64
	err := pgxscan.Get(ctx, d.pool, &count, sql, ids)
	if err != nil {
		return 0, err
	}

	return count, nil
}
