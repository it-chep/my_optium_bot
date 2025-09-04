package update_list_name

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/update_list_name/dal"
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

func (a *Action) Do(ctx context.Context, listID int64, name string) (err error) {
	return a.dal.UpdateListName(ctx, listID, name)
}
