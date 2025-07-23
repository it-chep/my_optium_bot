package init_chat

import (
	"context"
	"fmt"

	create_doctor_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/create_doctor/dal"
	init_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/doctors/init_chat/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/doctors/init_chat/service/create_patient"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/doctors/init_chat/service/step"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/doctors/init_chat/service/update_patient"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal        *create_doctor_dal.Dal
	common     *dal.CommonDal
	initBotDal *init_dal.InitBotDal
	bot        *tg_bot.Bot

	stepService    *step.Service
	patientCreator *create_patient.Service
	patientUpdater *update_patient.Service
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot, common *dal.CommonDal) *Action {
	return &Action{
		dal:        create_doctor_dal.NewDal(pool),
		common:     common,
		bot:        bot,
		initBotDal: init_dal.NewDal(pool),

		stepService:    step.NewService(common, bot),
		patientCreator: create_patient.NewService(init_dal.NewDal(pool)),
		patientUpdater: update_patient.NewService(init_dal.NewDal(pool)),
	}
}

// Handle разбивка на степы сценария "Старт"
func (a *Action) Handle(ctx context.Context, usr user.User, msg dto.Message) (err error) {
	// После выполнения шага происходит движение на шаг вперед и отправка сообщения
	defer func(err error) {
		a.stepService.MoveToNextStep(ctx, usr, msg, err)
	}(err)

	logger.Message(ctx, fmt.Sprintf("Хендлим шаги в сценарии init_chat, user: %d", usr.ID))
	//  Получаем пациента чтобы потом его обновить.
	patient, err := a.initBotDal.GetPatient(ctx, msg.ChatID)
	if err != nil {
		return err
	}
	// todo, такого вроде не должно возникнуть, но если вдруг
	if patient.IsEmpty() && usr.StepStat.StepOrder != 1 {
		return nil
	}

	// Получаем текущий шаг пользователя
	switch usr.StepStat.StepOrder {
	case 1:
		err = a.patientCreator.CreatePatient(ctx, msg)
	case 2:
		err = a.patientUpdater.UpdateSex(ctx, patient.ID, msg)
	case 3:
		err = a.patientUpdater.UpdateBirthDate(ctx, patient.ID, msg)
	case 4:
		err = a.initBotDal.UpdatePatientMetricsLink(ctx, patient.ID, msg.Text)
	}

	return err
}
