package therapy

import (
	"context"
	"fmt"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/samber/lo"
)

func (a *Action) Handle(ctx context.Context, usr user.User, msg dto.Message) error {
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
	nextStep, _ := scenario.StepByOrder(btn.NextStepOrder)
	if err = a.bot.SendMessage(bot_dto.Message{Chat: msg.ChatID, Text: nextStep.Text}); err != nil {
		return err
	}

	if nextStep.IsFinal {
		return a.common.CompleteScenario(ctx, patient.TgID, msg.ChatID)
	}

	return a.common.MoveStepPatient(ctx, patient.TgID, msg.ChatID, lo.FromPtr(nextStep.NextStep), lo.FromPtr(nextStep.NextDelay))
}
