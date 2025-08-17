package information

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Action Сценарий "Информация"
type Action struct {
	common *dal.CommonDal
	bot    *tg_bot.Bot
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot, common *dal.CommonDal) *Action {
	return &Action{
		common: common,
		bot:    bot,
	}
}
