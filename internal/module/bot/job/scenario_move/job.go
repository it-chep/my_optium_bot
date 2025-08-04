package scenario_move

import (
	"context"
	"fmt"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
)

// Job джоба для таких сценариев как обучение, когда нужно ждать некоторое время прежде чем отправить след мессадж
type Job struct {
	dal *dal.JobDal

	actions bot.JobActions
}

func NewJob(dal *dal.JobDal, actions bot.JobActions) *Job {
	return &Job{
		dal:     dal,
		actions: actions,
	}
}

func (j *Job) Do(ctx context.Context) {
	scenarios, err := j.dal.GetActiveScheduledScenarios(ctx)
	if err != nil {
		logger.Error(ctx, "не удалось получить активные сценарии для сдвига", err)
		return
	}

	if len(scenarios) == 0 {
		return
	}

	for _, scenario := range scenarios {
		action, ok := j.actions[scenario.ScenarioID]
		if !ok {
			continue
		}

		if err = action(ctx, scenario); err != nil {
			logger.Error(ctx, fmt.Sprintf("произошла ошибка продолжения сценария %d на шаг %d для пользователя %d: %s",
				scenario.ScenarioID, scenario.Step, scenario.PatientID, err.Error()), err)
			continue
		}
	}
	return
}
