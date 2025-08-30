package dal

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/create_information_post/dto"
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

func (r *Repository) CreateInformationPost(ctx context.Context, req dto.Request) error {
	sql := `
		insert into information_posts (name, posts_theme_id, order_in_theme, media_id, content_type_id, post_text)
		values ($1, $2, $3, $5, $6, $4)
	`

	args := []interface{}{
		req.PostName,
		req.ThemeID,
		req.Order,
		req.Message,
	}

	if req.ContentTypeID == 0 {
		args = append(args, nil, nil) // media id
	} else {
		args = append(args, req.MediaID, req.ContentTypeID)
	}

	_, err := r.pool.Exec(ctx, sql, args...)

	return err
}
