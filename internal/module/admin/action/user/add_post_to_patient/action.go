package add_post_to_patient

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/add_post_to_patient/dal"
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

func (a *Action) Do(ctx context.Context, patientID, postID int64) error {
	return a.dal.AddPostToPatient(ctx, patientID, postID)
}
