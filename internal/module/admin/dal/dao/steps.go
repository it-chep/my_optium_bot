package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
	"github.com/samber/lo"
)

type StepDao struct {
	xo.ScenarioStep
}

func (sd *StepDao) ToDomain() dto.Step {
	return dto.Step{
		ID:           int64(sd.ID),
		ScenarioID:   int64(sd.ScenarioID),
		ScenarioName: dto.ScenarioNameMap[int64(sd.ScenarioID)],
		StepOrder:    int64(sd.StepOrder),
		Content:      sd.Content,
		IsFinal:      sd.IsFinal.Bool,
		NextDelay:    lo.FromPtr(sd.NextDelay),
		NextStep:     sd.NextStep.Int64,
	}
}

type ListStepDao []StepDao

func (lsd ListStepDao) ToDomain() []dto.Step {
	domain := make([]dto.Step, 0, len(lsd))
	for _, step := range lsd {
		domain = append(domain, step.ToDomain())
	}
	return domain
}
