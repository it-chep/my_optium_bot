package dal

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/update_posts/dto"
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

func (r *Repository) UpdatePost(ctx context.Context, postID int64, post dto.Request) (err error) {
	sql := `
		update information_posts 
		set name=$1, 
		    posts_theme_id=$2,
		    order_in_theme=$3, 
		    media_id=$4, 
		    content_type_id=$5, 
		    post_text=$6
		where id = $7
	`
	args := []interface{}{
		post.PostName,
		post.ThemeID,
		post.Order,
		post.MediaID,
		post.ContentTypeID,
		post.Message,
		postID,
	}
	_, err = r.pool.Exec(ctx, sql, args...)
	return err
}
