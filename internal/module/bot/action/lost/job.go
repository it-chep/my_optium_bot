package lost

import (
	"context"
	"fmt"
	"time"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/template"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/samber/lo"
)

const (
	commDurationLost = 24 * time.Hour * 21
)

// Do джоба запускает через Do отложенные сообщения
func (a *Action) Do(ctx context.Context, ps dto.PatientScenario) error {
	patient, err := a.common.GetPatient(ctx, ps.PatientID)
	if err != nil {
		return err
	}

	if patient.LastCommunicate.After(time.Now().UTC().Add(-commDurationLost)) && ps.Step == 1 {
		_ = a.common.MoveScenario(ctx, ps.ID, patient.LastCommunicate.Add(commDurationLost).UTC())
		// костыль, чтобы не делать отдельное решение)
		return fmt.Errorf("не наступило время запуска сценария")
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
	})
}

type route struct {
	patient  user.Patient
	scenario dto.Scenario
	step     dto.Step
	ps       dto.PatientScenario
}

func (a *Action) route(ctx context.Context, r route) error {
	sendMSG := func() error {
		return a.bot.SendMessage(bot_dto.Message{
			Chat:    r.ps.ChatID,
			Text:    template.Execute(r.step.Text, r.patient),
			Buttons: r.step.Buttons,
		})
	}

	if err := sendMSG(); err != nil {
		return err
	}

	if r.step.NextStep != nil {
		if err := a.common.MoveStepPatient(ctx, dal.MoveStep{
			TgID:     r.patient.TgID,
			ChatID:   r.ps.ChatID,
			Scenario: r.scenario.ID,
			Step:     r.step.Order,
			NextStep: lo.FromPtr(r.step.NextStep),
			Delay:    lo.FromPtr(r.step.NextDelay),
			Answered: len(r.step.Buttons) == 0,
		}); err != nil {
			return err
		}

		if r.step.NextDelay == nil {
			return a.common.ScenarioNotAnswered(ctx, r.patient.TgID, r.ps.ScenarioID)
		}
	}

	if r.step.IsFinal {
		_ = a.common.CompleteScenario(ctx, r.patient.TgID, r.ps.ChatID, r.scenario.ID)
		next := time.Now().Add(7 * 24 * time.Hour)
		return a.common.AssignScenarios(ctx, r.patient.TgID, r.ps.ChatID, []dto.Scenario{{ID: 9, ScheduledTime: next}})
	}

	return a.common.MarkScenariosSent(ctx, r.ps)
}
