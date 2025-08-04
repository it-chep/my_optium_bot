package job

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/scenario_activate"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/scenario_move"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Aggregator struct {
	Activate *scenario_activate.Job
	Move     *scenario_move.Job
}

func NewAggregator(pool *pgxpool.Pool, actions bot.JobActions) *Aggregator {
	jobDal := dal.NewDal(pool)
	return &Aggregator{
		Activate: scenario_activate.NewJob(jobDal, actions),
		Move:     scenario_move.NewJob(jobDal, actions),
	}
}
