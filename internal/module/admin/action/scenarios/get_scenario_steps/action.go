package get_scenario_steps

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/get_scenario_steps/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Repository
}

func New(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, scenarioID int64) (_ []dto.Step, err error) {
	return a.dal.GetScenarioSteps(ctx, scenarioID)
}
