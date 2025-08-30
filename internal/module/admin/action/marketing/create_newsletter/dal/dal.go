package dal

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/create_newsletter/dto"
	indto "github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
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

func (d *Dal) CreateNewsletter(ctx context.Context, req dto.Request) error {
	sql := `
		insert into newsletters (
			recipients_count,
			text,
			users_lists,
			name,
			status_id,
			media_id,
			content_type_id
		)
		select 
			(select count(distinct user_id) from users_lists where list_id = any ($1::bigint[])),
			$2,
			$1::bigint[],
			$3,
			$4,
			$5,
			$6
	`
	args := []interface{}{
		req.UsersList,
		req.Text,
		req.Name,
		indto.Draft,
	}

	if req.ContentTypeID == 0 {
		args = append(args, nil, nil)
	} else {
		args = append(args, req.MediaID, req.ContentTypeID)
	}

	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}
