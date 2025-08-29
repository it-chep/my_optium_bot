package dao

import (
	"database/sql"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type InformationPost struct {
	xo.InformationPost
	ThemeIsRequired sql.NullBool `db:"theme_is_required"`
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
	}
}

func (ipl InformationPostList) ToDomain() []dto.InformationPost {
	domainPosts := make([]dto.InformationPost, 0, len(ipl))
	for _, post := range ipl {
		domainPosts = append(domainPosts, post.ToDomain())
	}
	return domainPosts
}
