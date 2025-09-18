package update_posts

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/update_posts/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/update_posts/dto"
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

func (a *Action) Do(ctx context.Context, postID int64, body dto.Request) error {
	return a.dal.UpdatePost(ctx, postID, body)
}
