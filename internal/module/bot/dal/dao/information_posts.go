package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/information"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type PostThemeDao struct {
	xo.PostsTheme
}

type InformationPostsDao struct {
	xo.InformationPost
}

func (ip InformationPostsDao) ToDomain() information.Post {
	return information.Post{
		ID:           ip.ID,
		Text:         ip.PostText,
		PostsThemeID: information.ThemeID(ip.PostsThemeID),
		OrderInTheme: int64(ip.OrderInTheme),
		MediaTgID:    ip.MediaID.String,
		Type:         dto.ContentType(ip.ContentTypeID.Int64),
	}
}
