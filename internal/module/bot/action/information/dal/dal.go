package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/information"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
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

// GetLastSentPost получаем последний отправленный пост, чтобы на основе него понять, какой нам надо получить
func (r *Repository) GetLastSentPost(ctx context.Context, patientID int64) (_ information.Post, err error) {
	lastSentPostSql := `
		select ip.* from information_posts ip
		left join patient_posts pp on ip.id = pp.post_id
		where pp.patient_id = $1 and pp.is_received is true
		order by pp.sent_time desc 
		limit 1
	`

	var lastSentPost dao.InformationPostsDao
	err = pgxscan.Get(ctx, r.pool, &lastSentPost, lastSentPostSql, patientID)
	if err != nil {
		// Если пользователю еще ничего не отправилось, то не ретурнем
		if !errors.Is(err, pgx.ErrNoRows) {
			return information.Post{}, err
		}
	}

	return lastSentPost.ToDomain(), nil
}

// GetNextPost получаем следующий не отправленный пост, по теме, которую определили на уровне бизнес логики
func (r *Repository) GetNextPost(ctx context.Context, patientID, themeID int64) (_ error) {
	nextPostsSql := `
		select ip.* from information_posts ip
		left join patient_posts pp on ip.id = pp.post_id
		where 
		    pp.patient_id = $1 and 
		    pp.is_received is false and 
		    ip.posts_theme_id = $2
		order by ip.order_in_theme asc 
		limit 1
	`

	var nextPosts dao.InformationPostsDao
	err := pgxscan.Get(ctx, r.pool, &nextPosts, nextPostsSql, patientID, themeID)
	if err != nil {
		// у пользователя нет назначенных информационных постов
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Message(ctx, "У пользователя нет назначенных постов")
			return nil
		}
		return err
	}

	return nil
}

// MarkPostSent помечает пост отправленным
func (r *Repository) MarkPostSent(ctx context.Context, patientID, postID int64) error {
	sql := `
		update patient_posts 
		set is_received = true, sent_time = now() 
		where patient_id = $1 and post_id = $2
	`

	_, err := r.pool.Exec(ctx, sql, patientID, postID)
	if err != nil {
		return err
	}
	return nil
}
