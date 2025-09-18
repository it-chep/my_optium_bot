package handler

import (
	"fmt"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	adminHandler "github.com/it-chep/my_optium_bot.git/internal/server/handler/admin"
	"net/http"

	"github.com/it-chep/my_optium_bot.git/internal/server/handler/auth"
	"github.com/it-chep/my_optium_bot.git/internal/server/middleware"

	"github.com/go-chi/chi/v5"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot"
)

type Config interface {
	Token() string
}

// TgHookParser .
type TgHookParser interface {
	HandleUpdate(r *http.Request) (*tgbotapi.Update, error)
}

type Handler struct {
	// todo: вероятно чето типо интерфейса будет с одним двумя методами (принимаем урл или контент тайп либо строку че тип ответил)
	botParser TgHookParser
	botModule *bot.Bot
	adminAgg  *adminHandler.HandlerAggregator
	router    *chi.Mux
}

func NewHandler(cfg Config, botParser TgHookParser, botModule *bot.Bot, adminModule *admin.Module) *Handler {
	h := &Handler{
		botParser: botParser,
		botModule: botModule,
		router:    chi.NewRouter(),
	}

	h.setupMiddleware()
	h.setupHandlerAggregator(adminModule)
	h.setupRoutes(cfg)

	return h
}

func (h *Handler) setupMiddleware() {
	h.router.Use(middleware.LoggerMiddleware)
	h.router.Use(middleware.CORSMiddleware)
}

func (h *Handler) setupHandlerAggregator(adminModule *admin.Module) {
	h.adminAgg = adminHandler.NewAggregator(adminModule)
}

func (h *Handler) setupRoutes(cfg Config) {
	h.router.Route("/", func(r chi.Router) {
		r.Post(fmt.Sprintf("/%s/", cfg.Token()), h.bot())
	})

	h.router.Route("/admin", func(r chi.Router) {
		r.Get("/", h.admin())

		// Auth routes
		r.Post("/login", auth.LoginHandler)
		r.Get("/check-token", auth.CheckValidHandler)
		r.Post("/test", middleware.JWTMiddleware(h.admin())) // example

		// Авторизация
		r.Post("/auth", h.adminAgg.Users.Auth.Handle())

		// Пользователи
		r.Route("/users", func(r chi.Router) {
			r.Get("/", h.adminAgg.Users.GetUsers.Handle())             // GET /admin/users
			r.Get("/{user_id}", h.adminAgg.Users.GetUserByID.Handle()) // GET /admin/users/{id}

			r.Post("/{user_id}/scheduled_time", h.adminAgg.Users.UpdateSheduleTime.Handle()) // POST /admin/users/{id}/scheduled_time

			r.Post("/{user_id}/post/{post_id}", h.adminAgg.Users.AddPostToPatient.Handle())        // POST /admin/users/{id}/post/{id}
			r.Delete("/{user_id}/post/{post_id}", h.adminAgg.Users.DeletePostFromPatient.Handle()) // DELETE /admin/users/{id}/post/{id}

			r.Post("/{user_id}/lists/{list_id}", h.adminAgg.Users.AddUserToList.Handle())        // POST /admin/users/{id}/lists/{id}
			r.Delete("/{user_id}/lists/{list_id}", h.adminAgg.Users.DeleteUserFromList.Handle()) // DELETE  /admin/users/{id}/lists/{id}
		})

		// Сценарии
		r.Route("/messages", func(r chi.Router) {
			r.Get("/", h.adminAgg.Scenarios.GetAdminMessages.Handle())                  // GET /admin/messages
			r.Post("/", h.adminAgg.Scenarios.CreateAdminMessage.Handle())               // POST /admin/messages
			r.Delete("/{message_id}", h.adminAgg.Scenarios.DeleteAdminMessage.Handle()) // DELETE /admin/messages/{id}
		})
		r.Route("/steps", func(r chi.Router) {
			r.Get("/", h.adminAgg.Scenarios.GetSteps.Handle())               // GET /admin/steps
			r.Post("/{step_id}", h.adminAgg.Scenarios.EditStepText.Handle()) // POST /admin/steps/{id}
		})
		r.Route("/scenarios", func(r chi.Router) {
			r.Get("/", h.adminAgg.Scenarios.GetScenarios.Handle())                        // GET /admin/scenarios
			r.Post("/{scenario_id}", h.adminAgg.Scenarios.EditScenarioDelay.Handle())     // POST /admin/scenarios/{id}
			r.Get("/{scenario_id}/steps", h.adminAgg.Scenarios.GetScenarioSteps.Handle()) // GET /admin/scenarios/{id}/steps
		})

		// Информационные посты
		r.Route("/information_posts", func(r chi.Router) {
			r.Get("/", h.adminAgg.InformationPost.GetInformationPosts.Handle())    // GET /admin/information_posts
			r.Get("/{post_id}", h.adminAgg.InformationPost.GetPostByID.Handle())   // GET /admin/information_posts/{id}
			r.Delete("/{post_id}", h.adminAgg.InformationPost.DeletePost.Handle()) // DELETE /admin/information_posts/{id}
			r.Post("/", h.adminAgg.InformationPost.CreateInformationPost.Handle()) // POST /admin/information_posts
			r.Post("/{post_id}", h.adminAgg.InformationPost.UpdatePost.Handle())   // POST /admin/information_posts/{id}
		})
		r.Route("/posts_themes", func(r chi.Router) {
			r.Get("/", h.adminAgg.InformationPost.GetPostsThemes.Handle())           // GET /admin/posts_themes
			r.Post("/", h.adminAgg.InformationPost.CreatePostTheme.Handle())         // POST /admin/posts_themes
			r.Delete("/{theme_id}", h.adminAgg.InformationPost.DeleteTheme.Handle()) // DELETE /admin/posts_themes/{id}
			r.Post("/{theme_id}", h.adminAgg.InformationPost.UpdateTheme.Handle())   // POST /admin/posts_themes/{id}
		})

		// Маркетинг
		r.Route("/newsletters", func(r chi.Router) {
			r.Get("/", h.adminAgg.Marketing.GetNewsLetters.Handle())                                    // GET /admin/newsletters
			r.Post("/", h.adminAgg.Marketing.CreateNewsletter.Handle())                                 // POST /admin/newsletters
			r.Post("/{newsletters_id}", h.adminAgg.Marketing.UpdateNewsletter.Handle())                 // POST /admin/newsletters/{id}
			r.Get("/{newsletters_id}", h.adminAgg.Marketing.GetNewsletterByID.Handle())                 // GET /admin/newsletters/{id}
			r.Delete("/{newsletters_id}", h.adminAgg.Marketing.DeleteNewsletter.Handle())               // DELETE /admin/newsletters/{id}
			r.Post("/{newsletters_id}/send_test_letter", h.adminAgg.Marketing.SendDraftLetter.Handle()) // POST /admin/newsletters/{id}/send_test_letter
			r.Post("/{newsletters_id}/send_letter", h.adminAgg.Marketing.SendLetterToUsers.Handle())    // POST /admin/newsletters/{id}/send_letter
		})
		r.Post("/recepients_count", h.adminAgg.Marketing.GetRecepientsCount.Handle()) // POST /admin/recepients_count
		r.Get("/content_types", h.adminAgg.Marketing.GetContentTypes.Handle())        // GET /admin/content_types
		r.Route("/users-lists", func(r chi.Router) {
			r.Get("/", h.adminAgg.Marketing.GetUsersLists.Handle())              // GET /admin/users-lists
			r.Post("/", h.adminAgg.Marketing.CreateUserList.Handle())            // POST /admin/users-lists
			r.Post("/{list_id}", h.adminAgg.Marketing.UpdateListName.Handle())   // POST /admin/users-lists/{id}
			r.Delete("/{list_id}", h.adminAgg.Marketing.DeleteUserList.Handle()) // DELETE /admin/users-lists/{id}
		})
	})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
