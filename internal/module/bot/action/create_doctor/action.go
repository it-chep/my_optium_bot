package create_doctor

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/create_doctor/create_doctor_dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
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

func (a *Action) CreateDoctor(ctx context.Context, msg dto.Message) error {
	user, err := a.bot.GetUser(msg)
	if err != nil {
		return err
	}

	if !user.IsAdmin {
		return a.bot.SendMessage(bot_dto.Message{
			Chat: msg.ChatID, Text: "Кажется, что вы не врач)",
		})
	}

	if err = a.dal.UpsertDoctor(ctx, user); err != nil {
		return err
	}

	scenario, err := a.common.GetScenario(ctx, 1)
	if err != nil {
		return err
	}

	if err = a.bot.SendMessage(bot_dto.Message{
		Chat: msg.ChatID, Text: "MyBot активирован. Здравствуйте, хозяин!",
	}); err != nil {
		return err
	}

	if err = a.bot.SendMessage(bot_dto.Message{
		Chat: msg.ChatID, Text: scenario.Steps[0].Text,
	}); err != nil {
		return err
	}

	// todo: доделать апдейт степа для врача

	return nil
}
