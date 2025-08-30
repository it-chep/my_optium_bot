package lost

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
)

// Action Сценарий "потеряшка"
type Action struct {
	common *dal.CommonDal
	bot    *tg_bot.Bot
}

func NewAction(common *dal.CommonDal, bot *tg_bot.Bot) *Action {
	return &Action{
		common: common,
		bot:    bot,
	}
}
