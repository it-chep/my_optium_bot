package exit

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/exit/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
)

type Action struct {
	bot  *tg_bot.Bot
	repo *dal.Dal
}

func NewAction(bot *tg_bot.Bot, repo *dal.Dal) *Action {
	return &Action{
		bot:  bot,
		repo: repo,
	}
}

func (a *Action) Do(ctx context.Context, msg dto.Message) error {
	err := a.repo.ExitScenario(ctx, msg.User)
	if err != nil {
		// todo log
		return err
	}
	return a.bot.SendMessage(bot_dto.Message{
		Chat: msg.ChatID, Text: "Выход из сценария",
	})
}
