package action

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/init"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Agg struct {
	Init *init.Action
}

func NewAgg(pool *pgxpool.Pool, bot *tg_bot.Bot) *Agg {
	return &Agg{
		Init: init.NewAction(pool, bot),
	}
}
