package init_bot

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/init_bot/init_dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *init_dal.Dal
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot) *Action {
	return &Action{
		dal: init_dal.NewDal(pool),
	}
}
