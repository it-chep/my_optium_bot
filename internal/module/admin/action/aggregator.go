package action

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/add_post_to_patient"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/add_user_to_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/auth"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/create_information_post"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/create_newsletter"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/create_post_theme"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/create_user_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/delete_post_from_patient"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/delete_user_from_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/delete_user_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/get_information_posts"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/get_posts_themes"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/get_users"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/get_users_lists"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Aggregator struct {
	// аутентификация
	Auth *auth.Action

	// списки пользователей
	AddUserToList      *add_user_to_list.Action
	DeleteUserFromList *delete_user_from_list.Action
	CreateUserList     *create_user_list.Action
	DeleteUserList     *delete_user_list.Action
	GetUsersLists      *get_users_lists.Action

	// рассылки
	CreateNewsLetter *create_newsletter.Action
	// SentDraftLetter
	// SendLetterToUsers

	// пользователи
	GetUsers *get_users.Action

	// сценарий информация
	GetPostsThemes        *get_posts_themes.Action
	GetInformationPosts   *get_information_posts.Action
	AddPostToPatient      *add_post_to_patient.Action
	CreateInformationPost *create_information_post.Action
	CreatePostTheme       *create_post_theme.Action
	DeletePostFromPatient *delete_post_from_patient.Action
}

func NewAggregator(pool *pgxpool.Pool) *Aggregator {
	return &Aggregator{
		Auth: auth.NewAction(),
		// списки пользователей
		AddUserToList:      add_user_to_list.NewAction(pool),
		DeleteUserFromList: delete_user_from_list.NewAction(pool),
		CreateUserList:     create_user_list.NewAction(pool),
		DeleteUserList:     delete_user_list.NewAction(pool),
		GetUsersLists:      get_users_lists.NewAction(pool),
		// рассылки
		CreateNewsLetter: create_newsletter.NewAction(),
		// пользователи
		GetUsers: get_users.NewAction(pool),
		// сценарий информация
		GetPostsThemes:        get_posts_themes.New(pool),
		GetInformationPosts:   get_information_posts.New(pool),
		AddPostToPatient:      add_post_to_patient.NewAction(pool),
		CreateInformationPost: create_information_post.New(pool),
		CreatePostTheme:       create_post_theme.NewAction(pool),
		DeletePostFromPatient: delete_post_from_patient.NewAction(pool),
	}
}
