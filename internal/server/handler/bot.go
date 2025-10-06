package handler

import (
	"github.com/samber/lo"
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
			if lo.Contains([]string{"left", "kicked"}, event.ChatMember.NewChatMember.Status) {
				return
			}
			usrID := event.ChatMember.NewChatMember.User.ID
			chat := event.ChatMember.Chat.ID
			_ = h.botModule.Actions.InvitePatient.InvitePatient(r.Context(), usrID, chat)
		}

		if !valid(event) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if event.FromChat() == nil || event.SentFrom() == nil {
			return
		}

		txt := ""
		mediaID := ""
		if event.Message != nil {
			txt = event.Message.Text
			// фото
			if event.Message.Photo != nil {
				// массив фото разбивает фотографию на 4 качества, берем самое плохое )
				mediaID = event.Message.Photo[0].FileID
			}
			// видео
			if event.Message.Video != nil {
				mediaID = event.Message.Video.FileID
			}
			// документ
			if event.Message.Document != nil {
				mediaID = event.Message.Document.FileID
			}
			// кружок
			if event.Message.VideoNote != nil {
				mediaID = event.Message.VideoNote.FileID
			}
			// голосовое сообщение
			if event.Message.Voice != nil {
				mediaID = event.Message.Voice.FileID
			}
			// аудио сообщение
			if event.Message.Audio != nil {
				mediaID = event.Message.Audio.FileID
			}
		} else if event.CallbackQuery != nil {
			txt = event.CallbackQuery.Data
		}

		msg := dto.Message{
			User:    event.SentFrom().ID,
			Text:    txt,
			ChatID:  event.FromChat().ID,
			MediaID: mediaID,
		}

		if err = h.botModule.Route(r.Context(), msg); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
	}
}
