package information

import (
	"context"
	"fmt"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/information"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/template"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
)

func (a *Action) sendRequiredPost(ctx context.Context, r route) (information.Post, error) {
	post, err := a.service.GetRequiredPost(ctx, r.ps.PatientID)
	if err != nil {
		return post, err
	}

	// постов нет для отправки
	if post.ID == 0 {
		logger.Message(ctx, fmt.Sprintf("для пользователя tg_id %d, нет постов для отправки", r.patient.TgID))
		return post, nil
	}

	err = a.sendPostToTg(ctx, r, post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (a *Action) sendMotivationPost(ctx context.Context, r route) (information.Post, error) {
	post, err := a.service.GetMotivationPost(ctx, r.ps.PatientID)
	if err != nil {
		return post, err
	}

	// постов нет для отправки
	if post.ID == 0 {
		logger.Message(ctx, fmt.Sprintf("для пользователя tg_id %d, нет постов для отправки", r.patient.TgID))
		return post, nil
	}

	err = a.sendPostToTg(ctx, r, post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (a *Action) sendSecondPartPost(ctx context.Context, r route) (information.Post, error) {
	post, err := a.service.GetSecondPartPost(ctx, r.ps.PatientID)
	if err != nil {
		return post, err
	}

	// постов нет для отправки
	if post.ID == 0 {
		logger.Message(ctx, fmt.Sprintf("для пользователя tg_id %d, нет постов для отправки", r.patient.TgID))
		return post, nil
	}

	err = a.sendPostToTg(ctx, r, post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (a *Action) sendAdditionalPost(ctx context.Context, r route) (information.Post, error) {
	post, err := a.service.GetAdditionalPost(ctx, r.ps.PatientID)
	if err != nil {
		return post, err
	}

	// постов нет для отправки
	if post.ID == 0 {
		logger.Message(ctx, fmt.Sprintf("для пользователя tg_id %d, нет постов для отправки", r.patient.TgID))
		return post, nil
	}

	err = a.sendPostToTg(ctx, r, post)
	if err != nil {
		return post, err
	}

	return post, nil
}

func (a *Action) sendPostToTg(ctx context.Context, r route, post information.Post) error {
	var (
		sentPostID int64
		err        error
	)

	sentPostID = post.ID
	tmpl := toTemplateStruct{
		RequiredPost: post.Text, AdditionalPost: post.Text, MotivationPost: post.Text, SecondPartPost: post.Text,
	}

	// Только если у контента есть ID из телеги мы отправляем это media
	if post.MediaTgID != "" {
		return a.bot.SendMessageWithContentType(bot_dto.Message{
			Chat:        r.msg.ChatID,
			MediaID:     post.MediaTgID,
			ContentType: post.Type,
			Text:        template.Execute(r.step.Text, tmpl),
		})
	}

	// Отправляем сообщение без медиа
	err = a.bot.SendMessage(bot_dto.Message{
		Chat:    r.msg.ChatID,
		Text:    template.Execute(r.step.Text, tmpl),
		Buttons: r.step.Buttons,
	})
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Ошибка при отправке поста %d, юзер: %d", sentPostID, r.msg.ChatID), err)
		return err
	}

	return nil
}
