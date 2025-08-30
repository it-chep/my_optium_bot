package edit_scenario_delay

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/edit_scenario_delay/dal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Dal
}

func NewAction(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewDal(pool),
	}
}

func (a *Action) Do(ctx context.Context, scenarioID int64, delay string) error {
	// todo валидацию на delay
	return a.dal.UpdateScenarioDelay(ctx, scenarioID, delay)
}
