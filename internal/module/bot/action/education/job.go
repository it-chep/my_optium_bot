package education

import (
	"context"
	"fmt"
	"strings"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/template"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/samber/lo"
)

type toTemplateStruct struct {
	user.Patient
	DoctorUsername string
}

// Do Сценарий "ОБУЧЕНИЕ"
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
	sendMSG := func() error {
		tmpl := toTemplateStruct{
			Patient: r.patient,
		}

		// Получаем видос из базы
		content, err := a.educationDal.GetStepContent(ctx, r.scenario.ID, int64(r.step.Order))
		if err != nil {
			logger.Error(ctx, "Ошибка при получении медиа из базы, ОБУЧЕНИЕ", err)
			return err
		}

		if strings.Contains(r.step.Text, "DoctorUsername") {
			tmpl.DoctorUsername = a.educationDal.GetDoctorUsername(ctx, r.msg.ChatID)
		}

		// Только если у контента есть ID из телеги мы отправляем это media
		if content.MediaTgID != "" {
			return a.bot.SendMessageWithContentType(bot_dto.Message{
				Chat:        r.msg.ChatID,
				MediaID:     content.MediaTgID,
				ContentType: content.Type,
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

	// Отправляем сообщение которое положено пользователю
	if err := sendMSG(); err != nil {
		return err
	}

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
