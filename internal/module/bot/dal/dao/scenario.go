package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
	"github.com/samber/lo"
)

type Scenario struct {
	xo.Scenario
}

func (s *Scenario) ToDomain(steps *Steps, buttons *Buttons) dto.Scenario {
	return dto.Scenario{
		ID:    int64(s.ID),
		Name:  dto.ScenarioName(s.Name),
		Steps: steps.ToDomain(buttons),
	}
}

type Steps []*Step

func (s Steps) ToDomain(buttons *Buttons) dto.Steps {
	return lo.Map(s, func(s *Step, _ int) dto.Step {
		step := dto.Step{
			ID:               int64(s.ID),
			ScenarioID:       int64(s.ScenarioID),
			Order:            s.StepOrder,
			Text:             s.Content,
			IsFinal:          s.IsFinal.Bool,
			NextDelay:        s.NextDelay,
			DelayFromPatient: s.DelayFromPatient.Bool,
			Buttons:          buttons.ByStep(s).ToDomain(),
		}
		if s.NextStep.Valid {
			step.NextStep = lo.ToPtr(int(s.NextStep.Int64))
		}
		return step
	})
}

type Step struct {
	xo.ScenarioStep
}

func (s *Step) ToDomain(buttons Buttons) dto.Step {
	return dto.Step{
		ID:               int64(s.ID),
		ScenarioID:       int64(s.ScenarioID),
		Order:            s.StepOrder,
		Text:             s.Content,
		IsFinal:          s.IsFinal.Bool,
		Buttons:          buttons.ToDomain(),
		DelayFromPatient: s.DelayFromPatient.Bool,
	}
}

type Buttons []*Button

func (b Buttons) ByStep(s *Step) Buttons {
	return lo.Filter(b, func(item *Button, _ int) bool {
		return item.Step == s.StepOrder && item.Scenario == s.ScenarioID
	})
}

func (b Buttons) ToDomain() dto.StepButtons {
	return lo.Map(b, func(but *Button, _ int) dto.StepButton {
		return dto.StepButton{
			Text:          but.ButtonText,
			NextStepOrder: int(but.NextStepOrder.Int64),
		}
	})
}

type Button struct {
	xo.StepButton
}
