package bot

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
)

func (b *Bot) Route(ctx context.Context, msg dto.Message) error {
	switch msg.Text {
	case "/init_bot":
		return b.Actions.CreateDoctor.CreateDoctor(ctx, msg)
	default:
		return b.routeScenario(ctx, msg)
	}
}

func (b *Bot) routeScenario(ctx context.Context, msg dto.Message) error {
	stat, err := b.commonDal.GetUser(ctx, msg.User, msg.ChatID)
	if err != nil {
		return err
	}

	if stat.StepStat == nil {
		return nil
	}

	action, ok := b.MessageActions[stat.StepStat.ScenarioID]
	if !ok {
		return nil
	}

	return action(ctx, stat, msg)
}
