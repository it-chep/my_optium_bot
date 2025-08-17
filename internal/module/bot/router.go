package bot

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
)

func (b *Bot) Route(ctx context.Context, msg dto.Message) error {

	switch msg.Text {
	case "/init_bot":
		return b.Actions.CreateDoctor.CreateDoctor(ctx, msg)
	case "/admin_exit":
		return b.Actions.Exit.Do(ctx, msg)
	case "/add_media":
		return b.Actions.AddMedia.Do(ctx, msg)
	case "/admin_chat":
		return b.Actions.CreateAdminChat.UpsertAdminChat(ctx, msg)
	default:
		return b.routeScenario(ctx, msg)
	}
}

func (b *Bot) routeScenario(ctx context.Context, msg dto.Message) error {
	if msg.MediaID != "" {
		doctor, err := b.commonDal.GetDoctorAddMedia(ctx, msg.ChatID)
		if err != nil {
			return err
		}
		if doctor.StepStat.StepOrder != 1 {
			return nil
		}
		return b.Actions.AddMedia.Do(ctx, msg)
	}

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
