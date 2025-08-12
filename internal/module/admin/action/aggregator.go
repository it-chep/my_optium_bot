package action

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/add_user_to_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/auth"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/create_newsletter"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/create_user_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/delete_user_from_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/delete_user_list"
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
}

func NewAggregator(pool *pgxpool.Pool) *Aggregator {
	return &Aggregator{
		Auth:               auth.NewAction(),
		AddUserToList:      add_user_to_list.NewAction(pool),
		DeleteUserFromList: delete_user_from_list.NewAction(pool),
		CreateUserList:     create_user_list.NewAction(pool),
		DeleteUserList:     delete_user_list.NewAction(pool),
		GetUsersLists:      get_users_lists.NewAction(pool),
		CreateNewsLetter:   create_newsletter.NewAction(),
		GetUsers:           get_users.NewAction(pool),
	}
}
