package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
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

func (d *Dal) GetNewsletter(ctx context.Context, letterID int64) (dto.Newsletter, error) {
	sql := `
		select * from newsletters where id = $1 and status_id = $2
	`

	args := []interface{}{
		letterID,
		dto.Draft,
	}

	var newsletter dao.Newsletter
	err := pgxscan.Get(ctx, d.pool, &newsletter, sql, args...)
	if err != nil {
		return dto.Newsletter{}, err
	}

	return newsletter.ToDomain(), nil
}
