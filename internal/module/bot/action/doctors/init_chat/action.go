package init_chat

import (
	"context"
	"fmt"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"time"

	create_doctor_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/create_doctor/dal"
	init_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/doctors/init_chat/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal        *create_doctor_dal.Dal
	common     *dal.CommonDal
	initBotDal *init_dal.InitBotDal
	bot        *tg_bot.Bot
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot, common *dal.CommonDal) *Action {
	return &Action{
		dal:    create_doctor_dal.NewDal(pool),
		common: common,
		bot:    bot,
	}
}

// Handle todo: тут уже делаем разбивку по степам конкретного сценария
func (a *Action) Handle(ctx context.Context, usr user.User, msg dto.Message) (err error) {
	// После выполнения шага происходит движение на шаг вперед и отправка сообщения
	defer func(err error) {
		if err != nil {
			return
		}
		a.moveToNextStep(ctx, usr, msg)
	}(err)

	logger.Message(ctx, fmt.Sprintf("Хендлим шаги в сценарии init_chat, user: %d", usr.ID))
	//  Получаем пациента чтобы потом его обновить.
	patient, err := a.initBotDal.GetPatient(ctx, usr.ID)
	//if err != nil {
	//	return err
	//}
	//if patient.IsEmpty() && usr.StepStat.StepOrder != 1 {
	//	return nil
	//}
	//  Если его нет и ошибка notFound и шаг не 1, то ретурн

	// Получаем текущий шаг пользователя
	switch usr.StepStat.StepOrder {
	case 1:
		// Устанавливаем фио пациента
		err = a.initBotDal.CreatePatient(ctx, msg.Text)
		if err != nil {
			return err
		}
	case 2:
		// Устанавливаем Пол пациента
		var sex user.Sex
		sex = user.Man
		if msg.Text == "Ж" {
			sex = user.Woman
		}

		err = a.initBotDal.UpdatePatientSex(ctx, sex)
		if err != nil {
			return err
		}
	case 3:
		// Устанавливаем дату рождения пациента
		var birthDate time.Time
		birthDate, err = time.Parse("02.01.2006", msg.Text)
		if err != nil {
			return err
		}

		err = a.initBotDal.UpdatePatientBirthDate(ctx, birthDate)
		if err != nil {
			return err
		}
	case 4:
		// Устанавливаем ссылку на метрики
		err = a.initBotDal.UpdatePatientMetricsLink(ctx, msg.Text)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *Action) moveToNextStep(ctx context.Context, usr user.User, msg dto.Message) {
	// todo проверка, точно ли мы должны подвинуть в стейте ?

	step, err := a.common.DoctorNextStep(ctx, usr)
	if err != nil {
		return
	}
	if step.Text == "" {
		return
	}

	message := bot_dto.Message{
		Chat: msg.ChatID, Text: step.Text,
	}
	err = a.bot.SendMessage(message)
	if err != nil {
		return
	}
}
