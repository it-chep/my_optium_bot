package handler

import (
	"fmt"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot"
)

// TgHookParser .
type TgHookParser interface {
	HandleUpdate(r *http.Request) (*tgbotapi.Update, error)
}

type Handler struct {
	// todo: вероятно чето типо интерфейса будет с одним двумя методами (принимаем урл или контент тайп либо строку че тип ответил)
	botParser TgHookParser
	botModule *bot.Bot
	// admin svc
}

func NewHandler(botParser TgHookParser, botModule *bot.Bot) *Handler {
	return &Handler{
		botParser: botParser,
		botModule: botModule,
	}
}

func (h Handler) HandleAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(fmt.Sprintf("произошло восстановление системы: %v", err))
			}
		}()

		switch r.URL.Path {
		case "/telegram-webhook":
			h.bot().ServeHTTP(w, r)
		case "/admin":
			h.admin().ServeHTTP(w, r)
		default:
			_, _ = w.Write([]byte("По-моему, ты перепутал"))
		}
	}
}
