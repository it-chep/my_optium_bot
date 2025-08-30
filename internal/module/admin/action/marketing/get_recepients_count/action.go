package get_recepients_count

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/get_recepients_count/dal"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Dal
}

func NewAction(pool *pgxpool.Pool) *Action {
	return &Action{dal: dal.New(pool)}
}

func (a *Action) Do(ctx context.Context, ids []int64) (int64, error) {
	return a.dal.GetRecipientsCount(ctx, ids)
}
