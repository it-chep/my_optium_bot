package handler

import (
	"fmt"
	"net/http"

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
	// admin svc
	router *chi.Mux
}

func NewHandler(cfg Config, botParser TgHookParser, botModule *bot.Bot) *Handler {
	h := &Handler{
		botParser: botParser,
		botModule: botModule,
		router:    chi.NewRouter(),
	}

	h.setupRoutes(cfg)

	return h
}

func (h *Handler) setupRoutes(cfg Config) {
	h.router.Route("/", func(r chi.Router) {
		r.Post(fmt.Sprintf("/%s/", cfg.Token()), h.bot())
	})

	h.router.Route("/admin", func(r chi.Router) {
		r.Get("/", h.admin())
	})
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}
