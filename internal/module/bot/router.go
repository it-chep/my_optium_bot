package bot

import (
	"context"
	"time"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/template"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
)

func (b *Bot) Route(ctx context.Context, msg dto.Message) error {
	switch msg.Text {
	case "/init_bot", "/init_chat":
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
		if b.commonDal.NeedToSendControl(ctx, msg.User, msg.ChatID) {
			return b.sendControl(ctx, msg)
		}
		return nil
	}

	action, ok := b.MessageActions[stat.StepStat.ScenarioID]
	if !ok {
		return nil
	}

	return action(ctx, stat, msg)
}

type ControlTpl struct {
	Text      string
	FullName  string
	BirthDate time.Time
}

func (b *Bot) sendControl(ctx context.Context, msg dto.Message) error {
	b.bot.SendMessage(bot_dto.Message{
		Chat: msg.ChatID,
		Text: "Спасибо, записал! Передам врачу и ассистенту клиники!",
	})

	patient, err := b.commonDal.GetPatient(ctx, msg.User)
	if err != nil {
		return err
	}

	tmpl := &ControlTpl{
		Text:      msg.Text,
		FullName:  patient.FullName,
		BirthDate: patient.BirthDate,
	}

	adminMessages, err := b.commonDal.GetAdminMessages(ctx, 10, 99)
	if err != nil {
		return err
	}
	for _, chatID := range adminMessages.ChatIDs {
		for _, message := range adminMessages.Messages {
			err = b.bot.SendMessage(bot_dto.Message{Chat: chatID, Text: template.Execute(message, tmpl)}, tg_bot.WithDisabledPreview())
			if err != nil {
				return err
			}
		}
	}

	doctorMessages, err := b.commonDal.GetDoctorMessages(ctx, msg.User, 10, 99)
	if err != nil {
		return err
	}
	for _, message := range doctorMessages.Messages {
		err = b.bot.SendMessage(bot_dto.Message{Chat: doctorMessages.DoctorID, Text: template.Execute(message, tmpl)}, tg_bot.WithDisabledPreview())
		if err != nil {
			return err
		}
	}
	return nil
}
