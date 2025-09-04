package dal

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/update_newsletter/dto"
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

func (r *Repository) UpdateNewsletter(ctx context.Context, newsletterID int64, body dto.UpdateNewsletterDTO) error {
	sql := `
		update newsletters
		set name            = $2,
			text            = $3,
			users_lists     = $4,
			media_id        = $5,
			content_type_id = $6
		where id = $1 
	`

	args := []interface{}{
		newsletterID,
		body.Name,
		body.Text,
		body.UsersLists,
		body.MediaID,
		body.ContentTypeID,
	}

	_, err := r.pool.Exec(ctx, sql, args...)
	return err
}
