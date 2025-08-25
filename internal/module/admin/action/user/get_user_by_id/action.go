package get_user_by_id

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/get_user_by_id/dal"
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
