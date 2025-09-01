package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) GetContentTypes(ctx context.Context) ([]dto.ContentTypeDTO, error) {
	sql := `select * from content_types`

	var types dao.ContentTypes
	err := pgxscan.Select(ctx, r.pool, &types, sql)
	if err != nil {
		return nil, err
	}

	return types.ToDomain(), err
}
