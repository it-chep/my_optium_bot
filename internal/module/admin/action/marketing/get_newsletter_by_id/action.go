package get_newsletter_by_id

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/get_newsletter_by_id/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Repository
}

func NewAction(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, letterID int64) (_ dto.Newsletter, err error) {
	return a.dal.GetNewsletter(ctx, letterID)
}
