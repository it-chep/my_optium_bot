package admin

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Module модуль отвечающий за работу админки
type Module struct {
	Actions *action.Aggregator
}

func New(pool *pgxpool.Pool, bot *tg_bot.Bot) *Module {
	actions := action.NewAggregator(pool, bot)

	return &Module{
		Actions: actions,
	}
}
