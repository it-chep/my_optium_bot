package init_chat

import (
	"context"

	create_doctor_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/create_doctor/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal    *create_doctor_dal.Dal
	common *dal.CommonDal
	bot    *tg_bot.Bot
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot, common *dal.CommonDal) *Action {
	return &Action{
		dal:    create_doctor_dal.NewDal(pool),
		common: common,
		bot:    bot,
	}
}

// Handle todo: тут уже делаем разбивку по степам конкретного сценария
func (a *Action) Handle(context.Context, user.User, dto.Message) error {
	return nil
}
