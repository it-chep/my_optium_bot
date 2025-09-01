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

func (r *Repository) GetPost(ctx context.Context, postID int64) (_ dto.InformationPost, err error) {
	sql := `
		select 
		    ip.*,
			ph.name           as posts_theme_name,
			ph.is_required    as theme_is_required
		from information_posts ip
        	join posts_themes ph on ip.posts_theme_id = ph.id
		where ip.id = $1
	`

	var posts dao.InformationPost
	err = pgxscan.Select(ctx, r.pool, &posts, sql, postID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.InformationPost{}, nil
		}
		return dto.InformationPost{}, err
	}

	return posts.ToDomain(), nil
}
