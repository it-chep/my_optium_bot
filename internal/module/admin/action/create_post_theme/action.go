package create_post_theme

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/create_post_theme/dal"
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

func (a *Action) Do(ctx context.Context, name string, isRequired bool) error {
	return a.dal.CreatePostTheme(ctx, name, isRequired)
}
