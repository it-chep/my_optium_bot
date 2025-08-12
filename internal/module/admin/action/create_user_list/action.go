package create_user_list

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/create_user_list/dal"
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

func (a *Action) Do(ctx context.Context, listName string) error {
	return a.dal.CreateUserList(ctx, listName)
}
