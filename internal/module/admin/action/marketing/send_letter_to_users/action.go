package send_letter_to_users

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/send_letter_to_users/dal"
	moduleBotDto "github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *dal.Dal
	bot *tg_bot.Bot
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot) *Action {
	return &Action{
		dal: dal.NewDal(pool),
		bot: bot,
	}
}

func (a *Action) Do(ctx context.Context, letterID int64) error {
	letter, err := a.dal.GetNewsletter(ctx, letterID)
	if err != nil {
		return err
	}

	users, err := a.dal.GetUsersToSend(ctx, letter.UsersLists)
	if err != nil {
		return err
	}

	// делаем фоновую отправку рассылки
	go func() {
		ctx = context.Background()

		_ = a.dal.MarkNewslettersIsPending(ctx, letterID)
		resultUserIDs := make([]int64, 0, len(users))

		for _, user := range users {
			msg := bot_dto.Message{
				MediaID:     letter.MediaID,
				Text:        letter.Text,
				ContentType: moduleBotDto.ContentType(letter.ContentType),
				Chat:        user.ChatID,
			}

			if letter.MediaID != "" {
				err = a.bot.SendMessageWithContentType(msg)
			} else {
				err = a.bot.SendMessage(msg)
			}
			if err != nil {
				continue
			}

			resultUserIDs = append(resultUserIDs, user.ID)
		}

		_ = a.dal.MarkNewslettersIsSent(ctx, letterID, resultUserIDs)
	}()

	return err
}
