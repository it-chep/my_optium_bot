package action

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/create_information_post"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/create_post_theme"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/delete_post"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/delete_post_theme"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/get_information_posts"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/get_post_by_id"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/get_posts_themes"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/update_post_theme"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/information_posts/update_posts"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/create_newsletter"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/create_user_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/delete_newsletter"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/delete_user_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/get_content_types"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/get_newsletter_by_id"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/get_newsletters"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/get_recepients_count"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/get_users_lists"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/send_draft_letter"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/send_letter_to_users"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/update_list_name"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/update_newsletter"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/create_admin_message"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/delete_admin_message"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/edit_scenario_delay"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/edit_step_text"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/get_admin_messages"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/get_scenario_steps"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/get_scenarios"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/scenarios/get_steps"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/add_post_to_patient"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/add_user_to_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/auth"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/delete_post_from_patient"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/delete_user_from_list"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/get_user_by_id"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/get_users"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/user/update_shedule_time"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Aggregator struct {
	// рассылки, списки пользователей
	CreateNewsLetter   *create_newsletter.Action
	CreateUserList     *create_user_list.Action
	DeleteUserList     *delete_user_list.Action
	GetUsersLists      *get_users_lists.Action
	GetNewsletterByID  *get_newsletter_by_id.Action
	GetNewsletters     *get_newsletters.Action
	GetRecepientsCount *get_recepients_count.Action
	SentDraftLetter    *send_draft_letter.Action
	SendLetterToUsers  *send_letter_to_users.Action
	GetContentTypes    *get_content_types.Action
	DeleteNewsteller   *delete_newsletter.Action
	UpdateNewsletter   *update_newsletter.Action
	UpdateListName     *update_list_name.Action

	// пользователи
	Auth                  *auth.Action
	GetUsers              *get_users.Action
	GetUserByID           *get_user_by_id.Action
	AddPostToPatient      *add_post_to_patient.Action
	DeletePostFromPatient *delete_post_from_patient.Action
	AddUserToList         *add_user_to_list.Action
	DeleteUserFromList    *delete_user_from_list.Action
	UpdateSheduleTime     *update_shedule_time.Action

	// сценарий информация
	GetPostsThemes        *get_posts_themes.Action
	GetInformationPosts   *get_information_posts.Action
	CreateInformationPost *create_information_post.Action
	CreatePostTheme       *create_post_theme.Action
	GetPostByID           *get_post_by_id.Action
	DeletePost            *delete_post.Action
	DeletePostTheme       *delete_post_theme.Action
	UpdateTheme           *update_post_theme.Action
	UpdatePost            *update_posts.Action

	// сценарии
	CreateAdminMessage *create_admin_message.Action
	DeleteAdminMessage *delete_admin_message.Action
	EditScenarioDelay  *edit_scenario_delay.Action
	EditStepText       *edit_step_text.Action
	GetScenarios       *get_scenarios.Action
	GetSteps           *get_steps.Action
	GetAdminMessages   *get_admin_messages.Action
	GetScenarioSteps   *get_scenario_steps.Action
}

func NewAggregator(pool *pgxpool.Pool, bot *tg_bot.Bot) *Aggregator {
	return &Aggregator{
		// рассылки, списки пользователей
		CreateUserList:     create_user_list.NewAction(pool),
		DeleteUserList:     delete_user_list.NewAction(pool),
		GetUsersLists:      get_users_lists.NewAction(pool),
		GetNewsletters:     get_newsletters.NewAction(pool),
		GetNewsletterByID:  get_newsletter_by_id.NewAction(pool),
		GetRecepientsCount: get_recepients_count.NewAction(pool),
		CreateNewsLetter:   create_newsletter.NewAction(pool),
		SentDraftLetter:    send_draft_letter.NewAction(pool, bot),
		SendLetterToUsers:  send_letter_to_users.NewAction(pool, bot),
		GetContentTypes:    get_content_types.New(pool),
		DeleteNewsteller:   delete_newsletter.New(pool),
		UpdateNewsletter:   update_newsletter.New(pool),
		UpdateListName:     update_list_name.New(pool),

		// пользователи
		GetUsers:              get_users.NewAction(pool),
		GetUserByID:           get_user_by_id.NewAction(pool),
		Auth:                  auth.NewAction(),
		AddPostToPatient:      add_post_to_patient.NewAction(pool),
		AddUserToList:         add_user_to_list.NewAction(pool),
		DeleteUserFromList:    delete_user_from_list.NewAction(pool),
		DeletePostFromPatient: delete_post_from_patient.NewAction(pool),
		UpdateSheduleTime:     update_shedule_time.NewAction(pool),

		// сценарий информация
		GetPostsThemes:        get_posts_themes.New(pool),
		GetInformationPosts:   get_information_posts.New(pool),
		CreateInformationPost: create_information_post.New(pool),
		CreatePostTheme:       create_post_theme.NewAction(pool),
		GetPostByID:           get_post_by_id.New(pool),
		DeletePost:            delete_post.New(pool),
		DeletePostTheme:       delete_post_theme.New(pool),
		UpdateTheme:           update_post_theme.New(pool),
		UpdatePost:            update_posts.New(pool),

		// сценарии
		CreateAdminMessage: create_admin_message.NewAction(pool),
		DeleteAdminMessage: delete_admin_message.NewAction(pool),
		EditScenarioDelay:  edit_scenario_delay.NewAction(pool),
		EditStepText:       edit_step_text.NewAction(pool),
		GetScenarios:       get_scenarios.NewAction(pool),
		GetSteps:           get_steps.NewAction(pool),
		GetAdminMessages:   get_admin_messages.NewAction(pool),
		GetScenarioSteps:   get_scenario_steps.New(pool),
	}
}
