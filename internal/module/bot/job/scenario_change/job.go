package scenario_change

import (
	"context"
	"time"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/job/dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
)

// Job джоба для таких сценариев как обучение, когда нужно ждать некоторое время прежде чем отправить след мессадж
type Job struct {
	dal *dal.JobDal
}

func NewJob(dal *dal.JobDal) *Job {
	return &Job{
		dal: dal,
	}
}

func (j *Job) Do(ctx context.Context) {
	scenarios, err := j.dal.GetActiveOldScenarios(ctx, time.Hour)
	if err != nil {
		logger.Error(ctx, "не удалось получить активные сценарии для изменения", err)
		return
	}

	if len(scenarios) == 0 {
		return
	}

	for _, scenario := range scenarios {
		if err = j.dal.MoveToFuture(ctx, scenario.ID, time.Now().UTC().Add(time.Hour)); err != nil {
			logger.Error(ctx, "не удалось подвинуть сценарий в будущее", err)
		}
	}

	return
}
