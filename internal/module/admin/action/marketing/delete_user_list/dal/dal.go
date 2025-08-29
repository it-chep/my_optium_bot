package dal

import (
	"context"
	"github.com/pkg/errors"

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

func (d *Dal) DeleteUserList(ctx context.Context, listID int64) error {
	sql := `
		delete from user_lists  where id = $1
		`
	result, err := d.pool.Exec(ctx, sql, listID)
	if result.RowsAffected() == 0 {
		return errors.New("Ошибка при удалении списка")
	}
	return err
}

func (d *Dal) DeleteUsersFromList(ctx context.Context, listID int64) error {
	sql := `
		delete from users_lists where list_id = $1
	`
	_, err := d.pool.Exec(ctx, sql, listID)
	return err
}
