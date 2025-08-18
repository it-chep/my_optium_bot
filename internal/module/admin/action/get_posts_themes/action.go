package get_posts_themes

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/get_posts_themes/dal"
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

func (a *Action) Do(ctx context.Context) (err error) {
	return a.dal.GetPostsThemes(ctx)
}
