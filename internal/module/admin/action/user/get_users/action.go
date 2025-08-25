package get_users

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/get_users/dal"

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

func (a *Action) Do(ctx context.Context) ([]dto.User, error) {
	return a.dal.GetUsers(ctx)
}
