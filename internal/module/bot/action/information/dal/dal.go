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
func (r *Repository) GetNextPost(ctx context.Context, patientID int64, lastSentPost information.Post) (post information.Post, err error) {
	nextPostsSql := `
		with additional_themes AS (
			select count(1) > 0 as has_any
			from information_posts ip
			join patient_posts pp on ip.id = pp.post_id
			where ip.posts_theme_id > 3
			and pp.patient_id = $1 
		)

		select ip.*, 
			(select has_any from additional_themes) as has_additional_themes
		from information_posts ip
				 join posts_themes ps on ip.posts_theme_id = ps.id
				 join patient_posts pp on ip.id = pp.post_id
				 join patients p on pp.patient_id = p.id
		where p.tg_id = $1
		  and pp.is_received is false
		  and ip.posts_theme_id != $2
		order by ps.theme_order desc,
				 ip.order_in_theme asc
	`

	var nextPosts []dao.InformationPostsDao
	err = pgxscan.Select(ctx, r.pool, &nextPosts, nextPostsSql, patientID, lastSentPost.PostsThemeID)
	if err != nil {
		// у пользователя нет назначенных информационных постов
		if errors.Is(err, pgx.ErrNoRows) {
			logger.Message(ctx, "У пользователя нет назначенных постов")
			return post, nil
		}
		return post, err
	}

	// получаем посты, которые еще не отправились, отсортированные сначала по теме, потом по номеру в теме
	for _, nextPost := range nextPosts {
		// берем первый пост
		if lastSentPost.ID == 0 {
			post = nextPost.ToDomain()
			break
		}
		// Если предыдущий был обязательной темой
		if lastSentPost.PostsThemeID == information.RequiredTheme {
			// Если у пользователя есть доп темы, то отправляем их
			if nextPost.HasAdditionalThemes && nextPost.PostsThemeID > 3 {
				post = nextPost.ToDomain()
				break
			}
			// Иначе отправляем мотивацию
			if nextPost.PostsThemeID == 2 {
				post = nextPost.ToDomain()
				break
			}
			continue
		}

		if lastSentPost.PostsThemeID == information.MotivationTheme {
			// если подошло время для "подготовки ко второму этапу" то отправляем его
			//if { //todo добавить join на пациентов запросить дату создания и по ней уже мерить + 2 месяца
			//	break
			//}
			// иначе отправляем "обязательную тему"
			if nextPost.PostsThemeID == 1 {
				post = nextPost.ToDomain()
				break
			}
			continue
		}
		// Если крайний был подготовкой ко второму этапу, то надо отправить "обязательную тему"
		if lastSentPost.PostsThemeID == information.PreparingToSecondTheme {
			if nextPost.PostsThemeID == 1 {
				post = nextPost.ToDomain()
				break
			}
			continue
		}

		// Если дошли сюда, то у нас предыдущий был дополнительный, значит надо отправить мотивацию
		if nextPost.PostsThemeID == 2 {
			post = nextPost.ToDomain()
			break
		}
	}

	return post, nil
}

// MarkPostSent помечает пост отправленным
func (r *Repository) MarkPostSent(ctx context.Context, patientID, postID int64) error {
	sql := `
		update patient_posts pp 
		set is_received = true, sent_time = now() 
		from patients p
		where pp.patient_id = p.id 
		  and p.tg_id = $1 
		  and pp.post_id = $2
	`

	_, err := r.pool.Exec(ctx, sql, patientID, postID)
	if err != nil {
		return err
	}
	return nil
}
