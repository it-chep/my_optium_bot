package handler

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
)

func valid(event *tgbotapi.Update) bool {
	if event.FromChat() == nil ||
		event.SentFrom() == nil {
		return false
	}

	return event.FromChat().IsGroup() || event.FromChat().IsSuperGroup()
}

func (h *Handler) bot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := h.botParser.HandleUpdate(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if event.ChatMember != nil {
			usr := event.ChatMember.NewChatMember.User.ID
			chat := event.ChatMember.Chat.ID
			if err = h.botModule.Actions.InvitePatient.InvitePatient(r.Context(), usr, chat); err != nil {
				return
			}
		}

		if !valid(event) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		msg := dto.Message{
			User:   event.SentFrom().ID,
			ChatID: event.FromChat().ID,
		}
		if event.Message != nil {
			msg.Text = event.Message.Text
		} else if event.CallbackQuery != nil {
			msg.Text = event.CallbackQuery.Data
		} else {
			return
		}
		if err = h.botModule.Route(r.Context(), msg); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
	}
}
