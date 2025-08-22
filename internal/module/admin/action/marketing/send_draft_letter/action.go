package send_draft_letter

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/send_draft_letter/dal"
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
