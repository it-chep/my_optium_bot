package get_newsletters

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/get_newsletters/dal"
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

func (a *Action) Do(ctx context.Context) (_ []dto.Newsletter, err error) {
	return a.dal.GetNewsletters(ctx)
}
