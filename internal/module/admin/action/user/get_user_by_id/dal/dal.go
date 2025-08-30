package dal

import (
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *Dal {
	return &Dal{
		pool: pool,
	}
}

func (d *Dal) GetUserByID(ctx context.Context, userID int64) (dto.User, error) {
	sql := `
		select * from patients where id = $1
	`

	var user dao.User
	if err := pgxscan.Get(ctx, d.pool, &user, sql, userID); err != nil {
		if pgxscan.NotFound(err) {
			return dto.User{}, errors.New("user not found")
		}
		return dto.User{}, err
	}

	return user.ToDomain(), nil
}

func (d *Dal) GetLists(ctx context.Context, userID int64) ([]dto.UsersList, error) {
	sql := `
		select ul.id, ul.name, 0
		from users_lists uls
				 join user_lists ul on uls.list_id = ul.id
		where uls.user_id = $1
	`

	var lists dao.UserLists
	if err := pgxscan.Select(ctx, d.pool, &lists, sql, userID); err != nil {
		return nil, fmt.Errorf("failed to get user lists: %w", err)
	}

	return lists.ToDomain(), nil
}

func (d *Dal) GetUnsentPosts(ctx context.Context, userID int64) ([]dto.InformationPost, error) {
	sql := `
	select ip.id,
		   ip.name,
		   pt.is_required as theme_is_required
	from patient_posts pp
			 join
		 information_posts ip on pp.post_id = ip.id
			 join
		 posts_themes pt on ip.posts_theme_id = pt.id
	where pp.patient_id = $1
	  and (pp.is_received is false or pp.is_received is null)
	order by pt.is_required desc, ip.name
	`

	var posts dao.InformationPostList
	if err := pgxscan.Select(ctx, d.pool, &posts, sql, userID); err != nil {
		return nil, err
	}

	return posts.ToDomain(), nil

}

func (d *Dal) GetScenarioInfo(ctx context.Context, userID int64) ([]dto.PatientScenario, error) {
	sql := `
		select * from patient_scenarios where patient_id = $1
	`

	var scenarios dao.PatientScenarios
	if err := pgxscan.Select(ctx, d.pool, &scenarios, sql, userID); err != nil {
		return nil, err
	}

	return scenarios.ToDomain(), nil
}
