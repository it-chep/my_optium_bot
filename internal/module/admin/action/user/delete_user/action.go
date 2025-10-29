package delete_user

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/delete_user/dal"
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

func (a *Action) Do(ctx context.Context, patientID int64) (err error) {
	tgID, err := a.dal.GetUserTgID(ctx, patientID)
	if err != nil {
		return err
	}
	err = a.dal.DeletePatientDoctor(ctx, patientID)
	if err != nil {
		return err
	}
	err = a.dal.DeletePatientPosts(ctx, tgID)
	if err != nil {
		return err
	}
	err = a.dal.DeletePatientScenarios(ctx, tgID)
	if err != nil {
		return err
	}
	err = a.dal.DeleteLists(ctx, patientID)
	if err != nil {
		return err
	}
	err = a.dal.DeleteUser(ctx, patientID)
	if err != nil {
		return err
	}
	err = a.dal.DeleteRepetitions(ctx, tgID)
	return err
}
