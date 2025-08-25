package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dal/dao"
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

func (r *Repository) GetInformationPosts(ctx context.Context) (err error) {
	sql := `
		select ip.*, ph.* from information_posts ip
		join posts_themes ph on ip.posts_theme_id = ph.id
	`

	var posts []dao.User // todo сделать посты
	err = pgxscan.Select(ctx, r.pool, &posts, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil
		}
		return err
	}

	return nil
}
