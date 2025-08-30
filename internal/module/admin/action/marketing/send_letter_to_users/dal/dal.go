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

// GetNewsletter получаем рассылку
func (d *Dal) GetNewsletter(ctx context.Context, letterID int64) (dto.Newsletter, error) {
	sql := `
		select * from newsletters where id = $1 and status_id = $2
	`

	args := []interface{}{
		letterID,
		dto.Draft,
	}

	var newsletter dao.Newsletter
	err := pgxscan.Get(ctx, d.pool, &newsletter, sql, args...)
	if err != nil {
		return dto.Newsletter{}, err
	}

	return newsletter.ToDomain(), nil
}

// GetUsersToSend получаем пользователей из списков рассылки
func (d *Dal) GetUsersToSend(ctx context.Context, listIDs []int64) ([]dto.User, error) {
	sql := `
		select p.* 
		from users_lists uls 
		    join patients p on uls.user_id = p.id 
		where uls.list_id = any ($1::bigint[])
	`

	var users dao.Users
	err := pgxscan.Select(ctx, d.pool, &users, sql, listIDs)
	if err != nil {
		return nil, err
	}

	return users.ToDomain(), nil
}

func (d *Dal) MarkNewslettersIsPending(ctx context.Context, letterID int64) error {
	sql := `
		update newsletters set status_id = $1 where id = $2
	`

	args := []interface{}{
		letterID,
		dto.Pending,
	}

	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}

func (d *Dal) MarkNewslettersIsSent(ctx context.Context, letterID int64, userIDs []int64) error {
	sql := `
		update newsletters set status_id = $1, sent_at = now(), users_ids = $2, recipients_count = $3 where id = $4
	`
	args := []interface{}{
		letterID,
		userIDs,
		len(userIDs),
		dto.Sent,
	}

	_, err := d.pool.Exec(ctx, sql, args...)
	return err

}
