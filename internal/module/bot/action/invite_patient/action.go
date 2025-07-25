package invite_patient

import (
	"context"

	invite_patient_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/invite_patient/dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *invite_patient_dal.Dal
	bot *tg_bot.Bot
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot) *Action {
	return &Action{
		dal: invite_patient_dal.NewDal(pool),
		bot: bot,
	}
}

func (a *Action) InvitePatient(ctx context.Context, tgID, chatID int64) error {
	if err := a.dal.SetUserTgID(ctx, tgID, chatID); err != nil {
		logger.Error(ctx, "ошибка финального создания юзера", err)
		return err
	}

	return a.bot.SendMessage(bot_dto.Message{
		Chat: chatID,
		Text: "TODO: тут назначение сценариев и инсерт в очередь",
	})
}
