package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) GetScenarioSteps(ctx context.Context, scenarioID int64) ([]dto.Step, error) {
	sql := `
		select * from scenario_steps
		where scenario_id = $1 
	`

	var steps dao.ListStepDao
	err := pgxscan.Select(ctx, r.pool, &steps, sql, scenarioID)
	if err != nil {
		return nil, err
	}
	return steps.ToDomain(), nil
}
