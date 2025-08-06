package bot

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/job_type"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageAction func(context.Context, user.User, dto.Message) error

type MessageActions map[int64]MessageAction

type Bot struct {
	Actions   *action.Agg
	commonDal *dal.CommonDal

	MessageActions MessageActions
	JobActions     job_type.JobActions

	Jobs *job.Aggregator
}

func New(pool *pgxpool.Pool, bot *tg_bot.Bot) *Bot {
	commonDal := dal.NewDal(pool)
	actions := action.NewAgg(pool, bot, commonDal)
	jobActions := job_type.JobActions{
		// такие сценарии как метрики (не нужно сообщение для триггера, периодичный сценарий)
		4: actions.Therapy.Do,
	}

	return &Bot{
		Actions:   actions,
		commonDal: commonDal,

		MessageActions: MessageActions{
			1: actions.InitChat.Handle,
			2: actions.Metrics.Handle,
			4: actions.Therapy.Handle,
		},
		JobActions: jobActions,
		Jobs:       job.NewAggregator(pool, jobActions),
	}
}
