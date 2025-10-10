package create_doctor

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"

	create_doctor_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/create_doctor/dal"
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

func (a *Action) CreateDoctor(ctx context.Context, msg dto.Message) (err error) {
	user, err := a.upsertDoctor(ctx, msg)
	if err != nil || !user.IsAdmin {
		logger.Error(ctx, "CreateDoctor err: %v", err)
		return err
	}

	scenario, err := a.common.GetScenario(ctx, 1)
	if err != nil {
		logger.Error(ctx, "GetScenario err: %v", err)
		return err
	}
	step := scenario.Steps[0]
	defer func() {
		if err == nil {
			err = a.common.UpdateDoctorStep(ctx, user.ID, step)
		}
	}()

	messages := []bot_dto.Message{
		{Chat: msg.ChatID, Text: "MyBot активирован. Здравствуйте, хозяин!"},
		{Chat: msg.ChatID, Text: step.Text},
	}
	return a.bot.SendMessages(messages)
}

func (a *Action) upsertDoctor(ctx context.Context, msg dto.Message) (bot_dto.User, error) {
	user, err := a.bot.GetUser(msg)
	if err != nil {
		return bot_dto.User{}, err
	}

	if !user.IsAdmin {
		return user, nil
		//return user, a.bot.SendMessage(bot_dto.Message{
		//	Chat: msg.ChatID, Text: "Кажется, что вы не врач)",
		//})
	}

	return user, a.dal.UpsertDoctor(ctx, user)
}

func (a *Action) setDoctorStep() {

}
