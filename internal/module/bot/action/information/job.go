package information

import (
	"context"
	"fmt"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/information"
	"strings"
	"time"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/template"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
	"github.com/samber/lo"
)

type toTemplateStruct struct {
	user.Patient
	RequiredPost, AdditionalPost, MotivationPost, SecondPartPost string
}

// Do Сценарий "Информации"
func (a *Action) Do(ctx context.Context, ps dto.PatientScenario) error {
	patient, err := a.common.GetPatient(ctx, ps.PatientID)
	if err != nil {
		return err
	}

	scenario, err := a.common.GetScenario(ctx, ps.ScenarioID)
	if err != nil {
		return err
	}

	step, ok := lo.Find(scenario.Steps, func(s dto.Step) bool {
		return ps.Step == s.Order
	})
	// todo: скорее всего овер фф, вероятно не понадобиться
	if !ok && ps.Step == 0 {
		step, ok = scenario.StepByOrder(1)
	}
	if !ok {
		return fmt.Errorf("step not found")
	}

	return a.routePost(ctx, route{
		patient:  patient,
		scenario: scenario,
		step:     step,
		ps:       ps,
		msg: dto.Message{
			User:   patient.TgID,
			ChatID: ps.ChatID,
		},
	})
}
func (a *Action) routePost(ctx context.Context, r route) error {
	// если надо отправить пост, то делаем отправку поста
	var (
		sentPostID int64
		post       information.Post
	)
	defer func() {
		err := a.service.MarkPostSent(ctx, r.patient.TgID, sentPostID)
		if err != nil {
			logger.Error(ctx, "Ошибка при отметке поста отправленным", err)
			return
		}
	}()

	if strings.Contains(r.step.Text, "RequiredPost") {
		shadowedPost, err := a.sendRequiredPost(ctx, r)
		sentPostID = shadowedPost.ID
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Ошибка при получение RequiredPost: %d", sentPostID), err)
			return err
		}
		post = shadowedPost
	} else if strings.Contains(r.step.Text, "MotivationPost") {
		shadowedPost, err := a.sendMotivationPost(ctx, r)
		sentPostID = shadowedPost.ID
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Ошибка при получение MotivationPost: %d", sentPostID), err)
			return err
		}
		post = shadowedPost
	} else if strings.Contains(r.step.Text, "SecondPartPost") {
		shadowedPost, err := a.sendSecondPartPost(ctx, r)
		sentPostID = shadowedPost.ID
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Ошибка при получение SecondPartPost: %d", sentPostID), err)
			return err
		}
		post = shadowedPost
	} else if strings.Contains(r.step.Text, "AdditionalPost") {
		shadowedPost, err := a.sendAdditionalPost(ctx, r)
		sentPostID = shadowedPost.ID
		if err != nil {
			logger.Error(ctx, fmt.Sprintf("Ошибка при получение AdditionalPost: %d", sentPostID), err)
			return err
		}
		post = shadowedPost
	} else {
		// Если нет, то отправляем простое сообщение
		if err := a.sendMsg(ctx, r); err != nil {
			return err
		}
		// Двигаем шаг пользователя
		if r.step.NextStep != nil {
			return a.common.MoveStepPatient(ctx, dal.MoveStep{
				TgID:     r.patient.TgID,
				ChatID:   r.msg.ChatID,
				Scenario: r.scenario.ID,
				Step:     r.step.Order,
				NextStep: lo.FromPtr(r.step.NextStep),
				Delay:    lo.FromPtr(r.step.NextDelay),
				Answered: len(r.step.Buttons) == 0,
			})
		}
	}

	if sentPostID != 0 {
		return a.movePatientStep(ctx, r, post)
	}

	// Завершаем сценарий
	if r.step.IsFinal {
		return a.common.CompleteScenario(ctx, r.patient.TgID, r.msg.ChatID, r.scenario.ID)
	}

	return a.common.MarkScenariosSent(ctx, r.ps)
}

// sendMsg отправка сообщения по сценарию
func (a *Action) sendMsg(_ context.Context, r route) error {
	return a.bot.SendMessage(bot_dto.Message{
		Chat:    r.ps.ChatID,
		Text:    template.Execute(r.step.Text, r.patient),
		Buttons: r.step.Buttons,
	})
}

func (a *Action) movePatientStep(ctx context.Context, r route, post information.Post) error {
	moveStep := dal.MoveStep{
		TgID:     r.patient.TgID,
		ChatID:   r.msg.ChatID,
		Scenario: r.scenario.ID,
		Step:     r.step.Order,
		Delay:    lo.FromPtr(r.step.NextDelay),
		Answered: len(r.step.Buttons) == 0,
	}

	count, _ := a.service.GetInformationScenRepetitionsNumber(ctx, r.patient.TgID)

	switch post.PostsThemeID {
	case information.RequiredTheme:
		// значит отправленный пост из обязательной темы
		additionalPosts, _ := a.service.GetAdditionalPost(ctx, r.patient.TgID)

		moveStep.NextStep = 5
		if additionalPosts.ID != 0 {
			// значит мы смогли получить пост и след степом отправим именно его
			moveStep.NextStep = 4
		}
	case information.MotivationTheme:
		// значит отправленный пост из темы мотивации
		moveStep.NextStep = 2
		if count >= 7 {
			moveStep.NextStep = 7
		} else {
			moveStep.Delay = 4 * 24 * time.Hour
			count++
			err := a.service.UpdateCounter(ctx, r.patient.TgID, count)
			if err != nil {
				logger.Error(ctx, "Ошибка при обновлении коунтера", err)
				return err
			}
		}
		_ = a.service.FinishScenarioOrContinue(ctx, r.patient.TgID)

	case information.PreparingToSecondTheme:
		// значит отправленный пост из подготовки ко 2 этапу
		moveStep.NextStep = 2
		moveStep.Delay = 4 * 24 * time.Hour
		_ = a.service.FinishScenarioOrContinue(ctx, r.patient.TgID)

		count++
		err := a.service.UpdateCounter(ctx, r.patient.TgID, count)
		if err != nil {
			logger.Error(ctx, "Ошибка при обновлении коунтера", err)
			return err
		}
	default:
		// значит пост из доп темы
		moveStep.NextStep = 5
	}

	return a.common.MoveStepPatient(ctx, moveStep)
}

//func (a *Action) route(ctx context.Context, r route) error {
//	// если надо отправить пост, то делаем отправку поста
//	if strings.Contains(r.step.Text, "InformationPost") {
//		if err := a.sendInformationPost(ctx, r); err != nil {
//			return err
//		}
//	} else {
//		// Если нет, то отправляем простое сообщение
//		if err := a.sendMsg(ctx, r); err != nil {
//			return err
//		}
//	}
//
//	_ = a.service.FinishScenarioOrContinue(ctx, r.patient.TgID)
//
//	// Двигаем шаг пользователя
//	if r.step.NextStep != nil {
//		return a.common.MoveStepPatient(ctx, dal.MoveStep{
//			TgID:     r.patient.TgID,
//			ChatID:   r.msg.ChatID,
//			Scenario: r.scenario.ID,
//			Step:     r.step.Order,
//			NextStep: lo.FromPtr(r.step.NextStep),
//			Delay:    lo.FromPtr(r.step.NextDelay),
//			Answered: len(r.step.Buttons) == 0,
//		})
//	}
//
//	// Завершаем сценарий
//	if r.step.IsFinal {
//		return a.common.CompleteScenario(ctx, r.patient.TgID, r.msg.ChatID, r.scenario.ID)
//	}
//
//	return a.common.MarkScenariosSent(ctx, r.ps)
//}

//func (a *Action) sendInformationPost(ctx context.Context, r route) error {
//	var sentPostID int64
//	defer func() {
//		err := a.service.MarkPostSent(ctx, r.patient.TgID, sentPostID)
//		if err != nil {
//			logger.Error(ctx, "Ошибка при отметке поста отправленным", err)
//			return
//		}
//	}()
//
//	post, err := a.service.GetNextPost(ctx, r.ps)
//	if err != nil {
//		return err
//	}
//	// постов нет для отправки
//	if post.ID == 0 {
//		logger.Message(ctx, fmt.Sprintf("для пользователя tg_id %d, нет постов для отправки", r.patient.TgID))
//		return nil
//	}
//
//	sentPostID = post.ID
//	tmpl := toTemplateStruct{
//		InformationPost: post.Text,
//	}
//
//	// Только если у контента есть ID из телеги мы отправляем это media
//	if post.MediaTgID != "" {
//		return a.bot.SendMessageWithContentType(bot_dto.Message{
//			Chat:        r.msg.ChatID,
//			MediaID:     post.MediaTgID,
//			ContentType: post.Type,
//			Text:        template.Execute(r.step.Text, tmpl),
//		})
//	}
//
//	// Отправляем сообщение без медиа
//	err = a.bot.SendMessage(bot_dto.Message{
//		Chat:    r.msg.ChatID,
//		Text:    template.Execute(r.step.Text, tmpl),
//		Buttons: r.step.Buttons,
//	})
//	if err != nil {
//		logger.Error(ctx, fmt.Sprintf("Ошибка при отправке поста %d, юзер: %d", sentPostID, r.msg.ChatID), err)
//		return err
//	}
//	return nil
//}
