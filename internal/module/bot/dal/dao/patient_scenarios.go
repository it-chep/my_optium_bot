package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
	"github.com/samber/lo"
)

type PatientScenarios []*PatientScenario

func (s PatientScenarios) ToDomain() []dto.PatientScenario {
	return lo.Map(s, func(sc *PatientScenario, _ int) dto.PatientScenario {
		return sc.ToDomain()
	})
}

type PatientScenario struct {
	xo.PatientScenario
}

func (s *PatientScenario) ToDomain() dto.PatientScenario {
	return dto.PatientScenario{
		ID:         int64(s.ID),
		ScenarioID: int64(s.ScenarioID),
		PatientID:  int64(s.PatientID),
		Step:       s.Step,
		ChatID:     int64(s.ChatID),
	}
}
