package marketing

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/create_newsletter"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/create_user_list"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/delete_user_list"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/get_content_types"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/get_newsletter_by_id"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/get_newsletters"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/get_recepients_count"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/get_users_lists"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/send_draft_letter"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing/send_letter_to_users"
)

type HandlerGroup struct {
	CreateNewsletter  *create_newsletter.Handler
	SendDraftLetter   *send_draft_letter.Handler
	SendLetterToUsers *send_letter_to_users.Handler

	GetUsersLists      *get_users_lists.Handler
	GetNewsLetters     *get_newsletters.Handler
	GetNewsletterByID  *get_newsletter_by_id.Handler
	GetRecepientsCount *get_recepients_count.Handler
	GetContentTypes    *get_content_types.Handler
	CreateUserList     *create_user_list.Handler
	DeleteUserList     *delete_user_list.Handler
}

func NewGroup(adminModule *admin.Module) *HandlerGroup {
	return &HandlerGroup{
		CreateNewsletter:  create_newsletter.NewHandler(adminModule),
		SendDraftLetter:   send_draft_letter.NewHandler(adminModule),
		SendLetterToUsers: send_letter_to_users.NewHandler(adminModule),

		GetUsersLists:      get_users_lists.NewHandler(adminModule),
		GetNewsLetters:     get_newsletters.NewHandler(adminModule),
		GetNewsletterByID:  get_newsletter_by_id.NewHandler(adminModule),
		GetContentTypes:    get_content_types.NewHandler(adminModule),
		GetRecepientsCount: get_recepients_count.NewHandler(adminModule),
		CreateUserList:     create_user_list.NewHandler(adminModule),
		DeleteUserList:     delete_user_list.NewHandler(adminModule),
	}
}
