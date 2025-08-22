package send_letter_to_users

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/send_letter_to_users/dal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Dal
}

// todo bot
func NewAction(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewDal(pool),
	}
}

// todo
