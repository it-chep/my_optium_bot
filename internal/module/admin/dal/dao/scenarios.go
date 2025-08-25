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
