package bot

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ScenarioAction func(context.Context, user.User, dto.Message) error

type Bot struct {
	Actions   *action.Agg
	commonDal *dal.CommonDal

	ScenarioActions map[int64]ScenarioAction
}

func New(pool *pgxpool.Pool, bot *tg_bot.Bot) *Bot {
	commonDal := dal.NewDal(pool)
	actions := action.NewAgg(pool, bot, commonDal)
	return &Bot{
		Actions:   actions,
		commonDal: commonDal,

		ScenarioActions: map[int64]ScenarioAction{
			1: actions.InitChat.Handle,
		},
	}
}
