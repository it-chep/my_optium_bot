package create_admin_chat

import (
	"context"

	create_admin_chat_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/create_admin_chat/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal *create_admin_chat_dal.Dal
	bot *tg_bot.Bot
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot) *Action {
	return &Action{
		dal: create_admin_chat_dal.NewDal(pool),
		bot: bot,
	}
}

func (a *Action) UpsertAdminChat(ctx context.Context, msg dto.Message) error {
	if err := a.dal.UpsertAdminChat(ctx, msg.ChatID); err != nil {
		return err
	}
	return a.bot.SendMessages([]bot_dto.Message{
		{
			Chat: msg.ChatID,
			Text: "Админ чат активирован, сюда будут пересылаться важные уведомления по прогрессу пациентов. Не выключайте уведомления",
		},
	})
}
