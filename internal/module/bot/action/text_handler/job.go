package text_handler

import (
	"context"
	"fmt"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/template"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/samber/lo"
)

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
		return a.common.MoveStepPatient(ctx, dal.MoveStep{
			TgID:     r.patient.TgID,
			ChatID:   r.ps.ChatID,
			Scenario: r.scenario.ID,
			Step:     r.step.Order,
			NextStep: lo.FromPtr(r.step.NextStep),
			Delay:    lo.FromPtr(r.step.NextDelay),
		})
	}

	if r.step.IsFinal {
		return a.common.CompleteScenario(ctx, r.patient.TgID, r.ps.ChatID, r.scenario.ID)
	}

	return a.common.MarkScenariosSent(ctx, r.ps)
}
