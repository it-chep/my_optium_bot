package add_media

import (
	"context"
	"fmt"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
)

type Action struct {
	bot    *tg_bot.Bot
	common *dal.CommonDal
}

func New(bot *tg_bot.Bot, common *dal.CommonDal) *Action {
	return &Action{
		bot:    bot,
		common: common,
	}
}

// Do Экшн обработки получения ID медиа файлов
func (a *Action) Do(ctx context.Context, msg dto.Message) error {
	scenario, err := a.common.GetScenario(ctx, 11)
	if err != nil {
		return err
	}

	step := scenario.Steps[0]
	if msg.Text == "/add_media" {
		err = a.common.UpdateDoctorStep(ctx, msg.User, dto.Step{ScenarioID: 11, Order: 1})
		if err != nil {
			return err
		}
		return a.bot.SendMessage(bot_dto.Message{
			Chat: msg.ChatID,
			Text: step.Text,
		})
	}

	return a.bot.SendMessage(bot_dto.Message{
		Chat: msg.ChatID,
		Text: fmt.Sprintf("ID вашего медиа файла \n\n%s", msg.MediaID),
	})
}
