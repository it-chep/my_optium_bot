package information_post

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/information_post/create_information_post"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/information_post/create_post_theme"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/information_post/get_information_posts"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/information_post/get_post_by_id"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/information_post/get_posts_themes"
)

type HandlerGroup struct {
	GetPostsThemes  *get_posts_themes.Handler
	CreatePostTheme *create_post_theme.Handler

	GetInformationPosts   *get_information_posts.Handler
	GetPostByID           *get_post_by_id.Handler
	CreateInformationPost *create_information_post.Handler
}

func NewGroup(adminModule *admin.Module) *HandlerGroup {
	return &HandlerGroup{
		GetPostsThemes:  get_posts_themes.NewHandler(adminModule),
		CreatePostTheme: create_post_theme.NewHandler(adminModule),

		GetInformationPosts:   get_information_posts.NewHandler(adminModule),
		GetPostByID:           get_post_by_id.NewHandler(adminModule),
		CreateInformationPost: create_information_post.NewHandler(adminModule),
	}
}
