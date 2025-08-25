package create_admin_message

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/create_admin_message/dal"
	actionDto "github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/create_admin_message/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Dal
}

func NewAction(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewDal(pool),
	}
}

func (a *Action) Do(ctx context.Context, request actionDto.CreateMessageRequest) error {
	if request.Type == dto.Admin {
		return a.dal.CreateAdminMessage(ctx, request.ScenarioID, request.ScenarioID, request.Message)
	}

	if request.Type == dto.Doctor {
		return a.dal.CreateDoctorMessage(ctx, request.ScenarioID, request.ScenarioID, request.Message)
	}

	// невалидный вопрос
	return nil
}
