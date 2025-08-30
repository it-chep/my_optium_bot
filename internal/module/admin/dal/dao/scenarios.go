package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
	"github.com/samber/lo"
)

type ScenarioDao struct {
	xo.Scenario
}

func (s *ScenarioDao) ToDomain() dto.Scenario {
	return dto.Scenario{
		ID:       int64(s.ID),
		Name:     dto.ScenarioNameMap[int64(s.ID)],
		IsActive: s.IsActive.Bool,
		Delay:    lo.FromPtr(s.Delay),
	}
}

type ListScenarioDao []ScenarioDao

func (lsd ListScenarioDao) ToDomain() []dto.Scenario {
	domain := make([]dto.Scenario, 0, len(lsd))
	for _, scen := range lsd {
		domain = append(domain, scen.ToDomain())
	}

	return domain
}

type PatientScenario struct {
	xo.PatientScenario
}

type PatientScenarios []PatientScenario

func (ps PatientScenario) ToDomain() dto.PatientScenario {
	return dto.PatientScenario{
		ID:            int64(ps.ID),
		PatientID:     int64(ps.PatientID),
		ChatID:        ps.ChatID,
		ScenarioID:    int64(ps.ScenarioID),
		Step:          int64(ps.Step),
		Answered:      ps.Answered,
		Sent:          ps.Sent,
		ScheduledTime: ps.ScheduledTime,
		Active:        ps.Active,
		Repeatable:    ps.Repeatable,
	}
}

func (psl PatientScenarios) ToDomain() []dto.PatientScenario {
	domain := make([]dto.PatientScenario, 0, len(psl))
	for _, ps := range psl {
		domain = append(domain, ps.ToDomain())
	}
	return domain
}
