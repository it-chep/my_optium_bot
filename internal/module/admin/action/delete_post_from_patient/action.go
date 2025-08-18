package delete_post_from_patient

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/delete_post_from_patient/dal"
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

func (a *Action) Do(ctx context.Context, patientID, postID int64) (err error) {
	return a.dal.DeletePostFromPatient(ctx, patientID, postID)
}
