package job_type

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
)

type JobAction func(context.Context, dto.PatientScenario) error

type JobActions map[int64]JobAction
