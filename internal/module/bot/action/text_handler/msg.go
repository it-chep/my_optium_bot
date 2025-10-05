package text_handler

import (
	"context"
	"fmt"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/template"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/samber/lo"
)

func (a *Action) Handle(ctx context.Context, usr user.User, msg dto.Message) (err error) {
	patient, err := a.common.GetPatient(ctx, msg.User)
	if err != nil {
		return err
	}

	scenario, err := a.common.GetScenario(ctx, usr.StepStat.ScenarioID)
	if err != nil {
		return err
	}

	step, ok := scenario.StepByOrder(int(usr.StepStat.StepOrder))
	if !ok {
		return fmt.Errorf("step not found")
	}

	btn, found := lo.Find(step.Buttons, func(b dto.StepButton) bool { return b.Text == msg.Text })
	if !found {
		return nil
	}

	defer func() {
		_ = a.common.UpdateLastCommunication(ctx, patient.ID)
	}()

	nextStep, _ := scenario.StepByOrder(btn.NextStepOrder)
	if err = a.bot.SendMessage(bot_dto.Message{Chat: msg.ChatID, Text: template.Execute(nextStep.Text, patient)}); err != nil {
		return err
	}

	defer func() {
		if err == nil {
			err = a.postAction(ctx, usr.StepStat.ScenarioID, int64(nextStep.Order), patient)
		}
	}()

	if nextStep.IsFinal {
		return a.common.CompleteScenario(ctx, patient.TgID, msg.ChatID, usr.StepStat.ScenarioID)
	}

	return a.common.MoveStepPatient(ctx,
		dal.MoveStep{
			TgID:     patient.TgID,
			ChatID:   msg.ChatID,
			Scenario: usr.StepStat.ScenarioID,
			Step:     int(usr.StepStat.StepOrder),
			NextStep: lo.FromPtr(nextStep.NextStep),
			Delay:    lo.FromPtr(nextStep.NextDelay),
			Answered: true,
		})
}

func (a *Action) postAction(ctx context.Context, scenario, step int64, patient user.Patient) error {
	adminMessages, err := a.common.GetAdminMessages(ctx, scenario, step)
	if err != nil {
		return err
	}
	for _, chatID := range adminMessages.ChatIDs {
		for _, message := range adminMessages.Messages {
			err = a.bot.SendMessage(bot_dto.Message{Chat: chatID, Text: template.Execute(message, patient)}, tg_bot.WithDisabledPreview())
			if err != nil {
				return err
			}
		}
	}

	doctorMessages, err := a.common.GetDoctorMessages(ctx, patient.TgID, scenario, step)
	if err != nil {
		return err
	}
	for _, message := range doctorMessages.Messages {
		err = a.bot.SendMessage(bot_dto.Message{Chat: doctorMessages.DoctorID, Text: template.Execute(message, patient)}, tg_bot.WithDisabledPreview())
		if err != nil {
			return err
		}
	}
	return nil
}
