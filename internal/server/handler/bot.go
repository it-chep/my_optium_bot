package handler

import (
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/samber/lo"
	"net/http"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
)

func (h *Handler) bot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		event, err := h.botParser.HandleUpdate(r)
		if err != nil {
			logger.Error(ctx, "Ошибка err: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if event.ChatMember != nil {
			if lo.Contains([]string{"left", "kicked"}, event.ChatMember.NewChatMember.Status) {
				return
			}
			usrID := event.ChatMember.NewChatMember.User.ID
			chat := event.ChatMember.Chat.ID
			err = h.botModule.Actions.InvitePatient.InvitePatient(ctx, usrID, chat)
			if err != nil {
				logger.Error(ctx, "Ошибка err: %v", err)
			}
		}

		if event.FromChat() == nil || event.SentFrom() == nil {
			logger.Message(ctx, "Невалидный хук")
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

		if err = h.botModule.Route(ctx, msg); err != nil {
			logger.Error(ctx, "ошибка при обработке ивента err: %v", err)
			return
		}
	}
}
