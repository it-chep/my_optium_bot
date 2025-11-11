package move_2_step

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/move_2_step/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/move_2_step/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/samber/lo"
	"time"
)

type Action struct {
	dal *dal.Dal
}

func New(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewDal(pool),
	}
}

func (a *Action) Do(ctx context.Context, userID int64) error {
	ids, err := a.dal.GetPatientIDs(ctx, userID)
	if err != nil {
		return err
	}

	patientScenarios, err := a.dal.GetPatientScenarios(ctx, ids)
	if err != nil {
		return err
	}

	scensMap := a.initScenarios()

	for _, scenario := range patientScenarios {
		newScen, ok := scensMap[scenario.ScenarioID]
		if !ok {
			continue
		}
		err := a.dal.MoveScenario(ctx, scenario.ID, newScen.ScheduledTime)
		if err != nil {
			return err
		}
	}
	return err
}

func (a *Action) initScenarios() map[int64]dto.Scenario {
	now := time.Now().UTC()

	// все сценарии будут начинаться в полдень по москве (чтобы не дудосить ночью пациентов)
	noon := time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, time.UTC)
	day := time.Hour * 24

	scenarios := []dto.Scenario{
		{ID: 4, ScheduledTime: noon.Add(1 * day)},   // терапия
		{ID: 6, ScheduledTime: noon.Add(2 * day)},   // рекомендации
		{ID: 2, ScheduledTime: noon.Add(3 * day)},   // метрики
		{ID: 9, ScheduledTime: noon.Add(21 * day)},  // потеряшка
		{ID: 10, ScheduledTime: noon.Add(60 * day)}, // выведение на контроль
	}

	return lo.SliceToMap(scenarios, func(scenario dto.Scenario) (int64, dto.Scenario) {
		return scenario.ID, scenario
	})
}
