package dao

import (
	"database/sql"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type InformationPost struct {
	xo.InformationPost
	ThemeIsRequired sql.NullBool   `db:"theme_is_required"`
	PostThemeName   sql.NullString `db:"posts_theme_name"`
}

type InformationPostList []InformationPost

func (ip InformationPost) ToDomain() dto.InformationPost {
	return dto.InformationPost{
		ID:              ip.ID,
		Name:            ip.Name,
		PostsThemeID:    ip.PostsThemeID,
		OrderInTheme:    ip.OrderInTheme,
		MediaID:         ip.MediaID.String,
		ContentTypeID:   ip.ContentTypeID.Int64,
		PostText:        ip.PostText,
		ThemeIsRequired: ip.ThemeIsRequired.Bool,
		ThemeName:       ip.PostThemeName.String,
	}
}

func (ipl InformationPostList) ToDomain() []dto.InformationPost {
	domainPosts := make([]dto.InformationPost, 0, len(ipl))
	for _, post := range ipl {
		domainPosts = append(domainPosts, post.ToDomain())
	}
	return domainPosts
}

type Theme struct {
	xo.PostsTheme
}

type ThemeList []Theme

func (t Theme) ToDomain() dto.PostTheme {
	return dto.PostTheme{
		ID:              t.ID,
		Name:            t.Name,
		ThemeIsRequired: t.IsRequired.Bool,
	}
}

func (tl ThemeList) ToDomain() []dto.PostTheme {
	domain := make([]dto.PostTheme, 0, len(tl))
	for _, t := range tl {
		domain = append(domain, t.ToDomain())
	}
	return domain
}

type InformationPostListView struct {
	ID   int64  `db:"id"`
	Name string `db:"name"`

	PostThemeName   string       `db:"posts_theme_name"`
	OrderInTheme    int          `db:"order_in_theme"`
	ThemeIsRequired sql.NullBool `db:"theme_is_required"`
}

type InformationPostListViewList []InformationPostListView

func (iplv InformationPostListViewList) ToDomain() []dto.InformationPostListView {
	domain := make([]dto.InformationPostListView, 0, len(iplv))
	for _, item := range iplv {
		domain = append(domain, dto.InformationPostListView{
			ID:              item.ID,
			Name:            item.Name,
			PostThemeName:   item.PostThemeName,
			OrderInTheme:    int64(item.OrderInTheme),
			ThemeIsRequired: item.ThemeIsRequired.Bool,
		})
	}

	return domain
}
