package create_newsletter

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/create_newsletter/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/create_newsletter/dto"
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

func (a *Action) Do(ctx context.Context, req dto.Request) error {
	return a.dal.CreateNewsletter(ctx, req)
}
