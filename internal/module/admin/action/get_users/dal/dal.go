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

func (d *Dal) GetUsers(ctx context.Context) ([]dto.User, error) {
	sql := `
		select p.*, 
		       uls.id as "lists.id",
               uls.name as "lists.name"
		from patients p
		    left join users_lists uls on p.id = uls.user_id
		    `

	var usersWithLists dao.Users
	err := pgxscan.Select(ctx, d.pool, &usersWithLists, sql)
	if err != nil {
		return nil, err
	}

	return usersWithLists.ToDomain(), nil
}
