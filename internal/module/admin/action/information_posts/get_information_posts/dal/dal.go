package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) GetInformationPosts(ctx context.Context) (_ []dto.InformationPostListView, err error) {
	sql := `
		select 
		    ip.id             as id,
			ip.name           as name,
			ph.name           as posts_theme_name,
			ip.order_in_theme as order_in_theme,
			ph.is_required    as theme_is_required
		from information_posts ip
        	join posts_themes ph on ip.posts_theme_id = ph.id
	`

	var posts dao.InformationPostListViewList
	err = pgxscan.Select(ctx, r.pool, &posts, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return posts.ToDomain(), nil
}
