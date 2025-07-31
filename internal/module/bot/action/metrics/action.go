package metrics

import (
	"context"
	"fmt"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
)

// Action Сценарий "Метрики"
type Action struct {
}

func NewAction() *Action {
	return &Action{}
}

// Handle разбивка на степы сценария "Метрики"
func (a *Action) Handle(ctx context.Context, usr user.User, msg dto.Message) (err error) {
	logger.Message(ctx, fmt.Sprintf("Хендлим шаги в сценарии Метрик, user: %d", usr.ID))

	// todo получаем пациента
	//patient, err := a.initBotDal.GetPatient(ctx, msg.ChatID)
	//if err != nil {
	//	return err
	//}
	//if patient.IsEmpty() && usr.StepStat.StepOrder != 1 {
	//	return nil
	//}

	// todo enum
	if usr.StepStat.ScenarioID == 1 {
		// Метрики начало
		// сообщение с "Да"/"Нет"
		// Развилка на "Да"/"Нет"

	}

	if usr.StepStat.ScenarioID == 2 {
		//Метрики повторение
		//
		//
		//Развилка на получилось ли заполнить метрики "Да"/"Нет"
	}

	return nil
}
