package scenario_move

import (
	"context"
	"fmt"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/job_type"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
)

// Job джоба для таких сценариев как обучение, когда нужно ждать некоторое время прежде чем отправить след мессадж
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
	scenarios, err := j.dal.GetActiveScheduledScenarios(ctx)
	if err != nil {
		logger.Error(ctx, "не удалось получить активные сценарии для сдвига", err)
		return
	}

	if len(scenarios) == 0 {
		return
	}

	sent := make([]dto.PatientScenario, 0, len(scenarios))
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
		sent = append(sent, scenario)
	}

	if err = j.dal.MarkScenariosSent(ctx, sent); err != nil {
		logger.Error(ctx, "не удалось пометить сценарии активными", err)
	}

	return
}
