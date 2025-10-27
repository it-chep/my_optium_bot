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
	"github.com/samber/lo"
	"sort"
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
		        join patients p on pp.patient_id = p.id
		where p.tg_id = $1 and pp.is_received is true
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
			join patients p on pp.patient_id = p.id
			where ip.posts_theme_id > 3
			and p.tg_id = $1
		)

		select ip.*, 
			(select has_any from additional_themes) as has_additional_themes
		from information_posts ip
				 join posts_themes ps on ip.posts_theme_id = ps.id
				 join patient_posts pp on ip.id = pp.post_id
		where pp.patient_id = $1 -- tg_id
		  and pp.is_received is false
		  and ip.posts_theme_id != $2
		order by ps.id,
				 ip.order_in_theme
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

	count, _ := r.GetSentPostsCount(ctx, patientID)

	themeMap := make(map[information.ThemeID][]dao.InformationPostsDao, 0)
	for _, nextPost := range nextPosts {
		themeMap[information.ThemeID(nextPost.PostsThemeID)] = append(themeMap[information.ThemeID(nextPost.PostsThemeID)], nextPost)
	}

	if lastSentPost.ID == 0 {
		return lo.FirstOrEmpty(themeMap[information.RequiredTheme]).ToDomain(), nil
	}

	if lastSentPost.PostsThemeID == information.RequiredTheme {
		additionalThemesIDs := lo.Filter(lo.Keys(themeMap), func(item information.ThemeID, _ int) bool {
			return !lo.Contains([]information.ThemeID{1, 2, 3}, item)
		})

		sort.Slice(additionalThemesIDs, func(i, j int) bool {
			return additionalThemesIDs[i] < additionalThemesIDs[j]
		})

		for _, themeID := range additionalThemesIDs {
			return lo.FirstOrEmpty(themeMap[themeID]).ToDomain(), nil
		}

		return lo.FirstOrEmpty(themeMap[information.MotivationTheme]).ToDomain(), nil
	}

	if lastSentPost.PostsThemeID == information.MotivationTheme {
		if count >= 8 {
			return lo.FirstOrEmpty(themeMap[information.PreparingToSecondTheme]).ToDomain(), nil
		}
		return lo.FirstOrEmpty(themeMap[information.RequiredTheme]).ToDomain(), nil
	}

	if lastSentPost.PostsThemeID == information.PreparingToSecondTheme {
		return lo.FirstOrEmpty(themeMap[information.RequiredTheme]).ToDomain(), nil
	}

	return lo.FirstOrEmpty(themeMap[information.MotivationTheme]).ToDomain(), nil
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

// GetSentPostsCount получение количества отправленных постов
func (r *Repository) GetSentPostsCount(ctx context.Context, patientID int64) (count int64, err error) {
	sql := `
		select count(*) from patient_posts pp join patients p on pp.patient_id = p.id where p.tg_id = $1 and pp.is_received is true
	`

	err = pgxscan.Get(ctx, r.pool, &count, sql, patientID)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FinishInformationScenario заканчивает информацию
func (r *Repository) FinishInformationScenario(ctx context.Context, patientID int64) error {
	sql := `
		update patient_scenarios set completed_at = now(), active = false, sent = true where patient_id = $1 and scenario_id = 8
	`

	_, err := r.pool.Exec(ctx, sql, patientID)
	return err
}
