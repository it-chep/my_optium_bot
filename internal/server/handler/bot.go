package handler

import (
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/samber/lo"
)

func isInit(event *tgbotapi.Update) bool {
	if event.Message != nil && len(event.Message.NewChatMembers) > 0 {
		// todo: написать ид бота
		return lo.ContainsBy(event.Message.NewChatMembers, func(usr tgbotapi.User) bool { return usr.ID == 1 })
	}
	return false
}

func valid(event *tgbotapi.Update) bool {
	if event.Message == nil ||
		event.FromChat() == nil ||
		event.SentFrom() == nil {
		return false
	}

	return (event.FromChat().IsGroup() || event.FromChat().IsSuperGroup()) &&
		event.Message.Text != ""
}

func (h Handler) bot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		event, err := h.botParser.HandleUpdate(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if isInit(event) {
			// todo: upsert врача, создание чата
		}

		if !valid(event) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err = h.botModule.Route(r.Context(), dto.Message{
			User: event.SentFrom().ID,
			Text: event.Message.Text,
		}); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
	}
}
