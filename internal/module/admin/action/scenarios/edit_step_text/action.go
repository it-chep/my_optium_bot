package edit_step_text

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/edit_step_text/dal"
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

func (a *Action) Do(ctx context.Context, stepID int64, text string) error {
	return a.dal.UpdateStepText(ctx, stepID, text)
}
