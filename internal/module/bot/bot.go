package bot

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessageAction func(context.Context, user.User, dto.Message) error

type MessageActions map[int64]MessageAction

type JobAction func(context.Context, dto.PatientScenario) error

type JobActions map[int64]JobAction

type Bot struct {
	Actions   *action.Agg
	commonDal *dal.CommonDal

	MessageActions MessageActions
	JobActions     JobActions

	Jobs *job.Aggregator
}

func New(pool *pgxpool.Pool, bot *tg_bot.Bot) *Bot {
	commonDal := dal.NewDal(pool)
	actions := action.NewAgg(pool, bot, commonDal)
	jobActions := JobActions{
		// такие сценарии как метрики (не нужно сообщение для триггера, периодичный сценарий)
	}

	return &Bot{
		Actions:   actions,
		commonDal: commonDal,

		MessageActions: MessageActions{
			1: actions.InitChat.Handle,
			2: actions.Metrics.Handle,
		},
		JobActions: jobActions,
		Jobs:       job.NewAggregator(pool, jobActions),
	}
}
