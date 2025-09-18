package delete_post_theme

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/delete_post_theme/dal"
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

func (a *Action) Do(ctx context.Context, themeID int64) error {
	err := a.dal.DeletePostTheme(ctx, themeID)
	if err != nil {
		return err
	}
	return a.dal.DropThemeFromPost(ctx, themeID)
}
