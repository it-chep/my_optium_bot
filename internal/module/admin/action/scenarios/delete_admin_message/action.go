package delete_admin_message

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/delete_admin_message/dal"
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

func (a *Action) Do(ctx context.Context, messageID int64, messageType dto.AdminType) error {
	if messageType == dto.Admin {
		return a.dal.DeleteAdminMessage(ctx, messageID)
	}

	if messageType == dto.Doctor {
		return a.dal.DeleteAdminMessage(ctx, messageID)
	}
	// невалидный запрос значит
	return nil
}
