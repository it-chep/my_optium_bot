package invite_patient

import (
	"context"
	"time"

	invite_patient_dal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/invite_patient/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Action struct {
	dal    *invite_patient_dal.Dal
	bot    *tg_bot.Bot
	common *dal.CommonDal
}

func NewAction(pool *pgxpool.Pool, bot *tg_bot.Bot, common *dal.CommonDal) *Action {
	return &Action{
		dal:    invite_patient_dal.NewDal(pool),
		bot:    bot,
		common: common,
	}
}

func (a *Action) InvitePatient(ctx context.Context, tgID, chatID int64) error {
	if err := a.dal.SetUserTgID(ctx, tgID, chatID); err != nil {
		logger.Error(ctx, "ошибка финального создания юзера", err)
		return err
	}

	return a.common.AssignScenarios(ctx, tgID, chatID, a.initScenarios())
}

func (a *Action) initScenarios() []dto.Scenario {
	now := time.Now().UTC()

	// все сценарии будут начинаться в полдень по москве (чтобы не дудосить ночью пациентов)
	noon := time.Date(now.Year(), now.Month(), now.Day(), 15, 0, 0, 0, time.UTC)
	day := time.Hour * 24

	return []dto.Scenario{
		// TODO: здесь будет инит очереди, по сути исерт начальных сценариев с соотв делеями
		{ID: 5, ScheduledTime: now.Add(10 * time.Minute)}, // обучение
		{ID: 4, ScheduledTime: noon.Add(1 * day)},         // терапия
		{ID: 6, ScheduledTime: noon.Add(2 * day)},         // рекомендации
		{ID: 2, ScheduledTime: noon.Add(4 * day)},         // метрики
		//{ID: 0, ScheduledTime: noon.Add(7 * day)},         // информация
		//{ID: 0, ScheduledTime: noon.Add(45 * day)},        // 2 этап
		{ID: 10, ScheduledTime: noon.Add(60 * day)}, // выведение на контроль
	}
}
