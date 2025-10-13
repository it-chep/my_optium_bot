package update_shedule_time

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/update_shedule_time/dal"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
	"time"
)

type Action struct {
	dal *dal.Dal
}

func NewAction(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewDal(pool),
	}
}

func (a *Action) Do(ctx context.Context, userID, scenarioID int64, nextDelay string) error {
	nextDelayTime, err := time.Parse(time.DateTime, nextDelay)
	if err != nil {
		return errors.New("Неправильно указан формат даты")
	}

	return a.dal.UpdateNextDelay(ctx, userID, scenarioID, nextDelayTime.Add(-3*time.Hour))
}
