package get_admin_messages

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/get_admin_messages/dal"
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

func (a *Action) Do(ctx context.Context) ([]dto.AdminMessage, error) {
	return a.dal.GetMessages(ctx)
}
