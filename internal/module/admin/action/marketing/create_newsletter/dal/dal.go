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

// todo
func (d *Dal) CreateNewsletter(ctx context.Context, userID, listID int64) error {
	return nil
}
