package delete_post

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/delete_post/dal"
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

func (a *Action) Do(ctx context.Context, postID int64) (err error) {
	err = a.dal.DeletePost(ctx, postID)
	if err != nil {
		return err
	}
	return a.dal.DeleteUnsentPost(ctx, postID)
}
