package information

import (
	"context"
	"fmt"
	"strings"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/template"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/samber/lo"
)

type toTemplateStruct struct {
	user.Patient
	InformationPost string
}

// Do Сценарий "Информации"
func (a *Action) Do(ctx context.Context, ps dto.PatientScenario) error {
	patient, err := a.common.GetPatient(ctx, ps.PatientID)
	if err != nil {
		return err
	}

	scenario, err := a.common.GetScenario(ctx, ps.ScenarioID)
	if err != nil {
		return err
	}

	step, ok := lo.Find(scenario.Steps, func(s dto.Step) bool {
		return ps.Step == s.Order
	})
	// todo: скорее всего овер фф, вероятно не понадобиться
	if !ok && ps.Step == 0 {
		step, ok = scenario.StepByOrder(1)
	}
	if !ok {
		return fmt.Errorf("step not found")
	}

	return a.route(ctx, route{
		patient:  patient,
		scenario: scenario,
		step:     step,
		ps:       ps,
		msg: dto.Message{
			User:   patient.TgID,
			ChatID: ps.ChatID,
		},
	})
}

func (a *Action) route(ctx context.Context, r route) error {
	// если надо отправить пост, то делаем отправку поста
	if strings.Contains(r.step.Text, "InformationPost") {
		if err := a.sendInformationPost(ctx, r); err != nil {
			return err
		}
	} else {
		// Если нет, то отправляем простое сообщение
		if err := a.sendMsg(ctx, r); err != nil {
			return err
		}
	}

	_ = a.service.FinishScenarioOrContinue(ctx, r.patient.TgID)

	// Двигаем шаг пользователя
	if r.step.NextStep != nil {
		return a.common.MoveStepPatient(ctx, dal.MoveStep{
			TgID:     r.patient.TgID,
			ChatID:   r.msg.ChatID,
			Scenario: r.scenario.ID,
			Step:     r.step.Order,
			NextStep: lo.FromPtr(r.step.NextStep),
			Delay:    lo.FromPtr(r.step.NextDelay),
			Answered: len(r.step.Buttons) == 0,
		})
	}

	// Завершаем сценарий
	if r.step.IsFinal {
		return a.common.CompleteScenario(ctx, r.patient.TgID, r.msg.ChatID, r.scenario.ID)
	}

	return a.common.MarkScenariosSent(ctx, r.ps)
}

// sendMsg отправка сообщения по сценарию
func (a *Action) sendMsg(_ context.Context, r route) error {
	return a.bot.SendMessage(bot_dto.Message{
		Chat:    r.ps.ChatID,
		Text:    template.Execute(r.step.Text, r.patient),
		Buttons: r.step.Buttons,
	})
}

func (a *Action) sendInformationPost(ctx context.Context, r route) error {
	var sentPostID int64
	defer func() {
		err := a.service.MarkPostSent(ctx, r.patient.TgID, sentPostID)
		if err != nil {
			logger.Error(ctx, "Ошибка при отметке поста отправленным", err)
			return
		}
	}()

	post, err := a.service.GetNextPost(ctx, r.ps)
	if err != nil {
		return err
	}
	// постов нет для отправки
	if post.ID == 0 {
		logger.Message(ctx, fmt.Sprintf("для пользователя tg_id %d, нет постов для отправки", r.patient.TgID))
		return nil
	}

	sentPostID = post.ID
	tmpl := toTemplateStruct{
		InformationPost: post.Text,
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
	return a.bot.SendMessage(bot_dto.Message{
		Chat:    r.msg.ChatID,
		Text:    template.Execute(r.step.Text, tmpl),
		Buttons: r.step.Buttons,
	})
}
