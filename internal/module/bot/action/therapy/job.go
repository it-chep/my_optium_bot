package therapy

import (
	"context"
	"fmt"

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

func (a *Action) route(_ context.Context, r route) error {
	sendMSG := func() error {
		return a.bot.SendMessage(bot_dto.Message{
			Chat:    r.ps.ChatID,
			Text:    template.Execute(r.step.Text, r.patient),
			Buttons: r.step.Buttons,
		})
	}

	return sendMSG()
}
