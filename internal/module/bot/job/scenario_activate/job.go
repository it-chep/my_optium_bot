package scenario_activate

import (
	"context"
	"fmt"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/job_type"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
)

// Job джоба для активации сценариев, проверяет не наступило ли время для триггера
type Job struct {
	dal *dal.JobDal

	actions job_type.JobActions
}

func NewJob(dal *dal.JobDal, actions job_type.JobActions) *Job {
	return &Job{
		dal:     dal,
		actions: actions,
	}
}

func (j *Job) Do(ctx context.Context) {
	scenarios, err := j.dal.GetAvailableScenarios(ctx)
	if err != nil {
		logger.Error(ctx, "ошибка получения доступных для запуска сценариев", err)
		return
	}

	if len(scenarios) == 0 {
		return
	}

	active := make([]dto.PatientScenario, 0, len(scenarios))
	for _, scenario := range scenarios {
		action, ok := j.actions[scenario.ScenarioID]
		if !ok {
			continue
		}

		if err = action(ctx, scenario); err != nil {
			logger.Error(ctx, fmt.Sprintf("произошла ошибка запуска сценария %d для пользователя %d: %s",
				scenario.ScenarioID, scenario.PatientID, err.Error()), err)
			continue
		}
		active = append(active, scenario)
	}

	if err = j.dal.MarkScenariosActive(ctx, active); err != nil {
		logger.Error(ctx, "не удалось пометить сценарии активными", err)
	}
}
