package get_user_by_id

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/get_user_by_id/dal"
	actionDto "github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/get_user_by_id/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type Action struct {
	dal *dal.Dal
}

func NewAction(pool *pgxpool.Pool) *Action {
	return &Action{
		dal: dal.NewDal(pool),
	}
}

func (a *Action) Do(ctx context.Context, patientID int64) (actionDto.Response, error) {
	var (
		wg          sync.WaitGroup
		userData    dto.User
		lists       []dto.UsersList
		scenarios   []dto.PatientScenario
		unsentPosts []dto.InformationPost
		userError   error
	)

	wg.Add(4)

	go func() {
		defer wg.Done()
		userData, userError = a.dal.GetUserByID(ctx, patientID)
	}()

	go func() {
		defer wg.Done()
		lists, _ = a.dal.GetLists(ctx, patientID)
	}()

	go func() {
		defer wg.Done()
		scenarios, _ = a.dal.GetScenarioInfo(ctx, patientID)
	}()

	go func() {
		defer wg.Done()
		unsentPosts, _ = a.dal.GetUnsentPosts(ctx, patientID)
	}()

	wg.Wait()

	return actionDto.Response{
		UserData:  userData,
		Scenarios: scenarios,
		Posts:     unsentPosts,
		Lists:     lists,
	}, userError
}
