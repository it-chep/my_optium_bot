package dal

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *Dal {
	return &Dal{
		pool: pool,
	}
}

func (d *Dal) GetUsersLists(ctx context.Context) ([]dto.UsersList, error) {
	sql := `
			select ul.id, ul.name, count(uls.*) as users_count 
			from user_lists ul
			left join users_lists uls on ul.id = uls.list_id 
			group by ul.id
		`

	var userLists dao.UserLists
	err := pgxscan.Select(ctx, d.pool, &userLists, sql)
	if err != nil {
		return nil, err
	}

	return userLists.ToDomain(), nil
}
