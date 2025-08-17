package add_user_to_list

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/add_user_to_list/dal"
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

func (a *Action) Do(ctx context.Context, userID, listID int64) error {
	return a.dal.CreateUserList(ctx, userID, listID)
}
