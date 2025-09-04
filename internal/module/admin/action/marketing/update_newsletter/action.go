package update_newsletter

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/update_newsletter/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/update_newsletter/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Repository
}

func New(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewRepository(pool),
	}
}

func (a *Action) Do(ctx context.Context, newsletterID int64, body dto.UpdateNewsletterDTO) (err error) {
	return a.dal.UpdateNewsletter(ctx, newsletterID, body)
}
