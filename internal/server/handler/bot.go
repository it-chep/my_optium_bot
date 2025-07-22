package handler

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
)

func valid(event *tgbotapi.Update) bool {
	if event.Message == nil ||
		event.FromChat() == nil ||
		event.SentFrom() == nil {
		return false
	}

	return (event.FromChat().IsGroup() || event.FromChat().IsSuperGroup()) &&
		event.Message.Text != ""
}

func (h *Handler) bot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := h.botParser.HandleUpdate(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !valid(event) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		msg := dto.Message{
			User:   event.SentFrom().ID,
			Text:   event.Message.Text,
			ChatID: event.FromChat().ID,
		}
		if err = h.botModule.Route(r.Context(), msg); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
	}
}
