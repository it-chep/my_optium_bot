package get_post_by_id

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/get_post_by_id/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
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

func (a *Action) Do(ctx context.Context, postID int64) (_ dto.InformationPost, err error) {
	return a.dal.GetPost(ctx, postID)
}
