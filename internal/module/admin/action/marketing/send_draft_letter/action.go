package send_draft_letter

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/action/marketing/send_draft_letter/dal"
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

	msg := bot_dto.Message{
		MediaID:     letter.MediaID,
		Text:        letter.Text,
		ContentType: moduleBotDto.ContentType(letter.ContentType),
		Chat:        0, // предлагаю захардкодить
	}

	if letter.MediaID != "" {
		err = a.bot.SendMessageWithContentType(msg)
	} else {
		err = a.bot.SendMessage(msg)
	}

	return err
}
