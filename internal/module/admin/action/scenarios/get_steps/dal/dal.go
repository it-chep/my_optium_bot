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

func (d *Dal) GetSteps(ctx context.Context) ([]dto.Step, error) {
	sql := `
		select st.* from scenario_steps st
	`

	var stepsDao dao.ListStepDao
	err := pgxscan.Select(ctx, d.pool, &stepsDao, sql)
	if err != nil {
		return nil, err
	}

	return stepsDao.ToDomain(), nil
}
