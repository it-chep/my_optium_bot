package step

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/logger"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot/bot_dto"
)

type Dal interface {
	DoctorNextStep(ctx context.Context, usr user.User, chatID int64) (dto.Step, error)
}

type Service struct {
	stepDal Dal
	bot     *tg_bot.Bot
}

func NewService(stepDal Dal, bot *tg_bot.Bot) *Service {
	return &Service{
		stepDal: stepDal,
		bot:     bot,
	}
}

func (s *Service) MoveToNextStep(ctx context.Context, usr user.User, msg dto.Message, err error) {
	if err != nil {
		// todo сделать кастомный текст для ошибки типа "вы вели не валидную дату рождения"
		s.sendErrorMessage(ctx, msg, "Что-то пошло не так, не смог продвинуть вас дальше по сценарию")
		return
	}
	s.moveToNextStep(ctx, usr, msg)
}

func (s *Service) moveToNextStep(ctx context.Context, usr user.User, msg dto.Message) {
	// todo проверка, точно ли мы должны подвинуть в стейте ?

	step, err := s.stepDal.DoctorNextStep(ctx, usr, msg.ChatID)
	if err != nil {
		return
	}
	if step.Text == "" {
		return
	}

	message := bot_dto.Message{
		Chat: msg.ChatID, Text: step.Text, Buttons: step.Buttons,
	}
	err = s.bot.SendMessage(message)
	if err != nil {
		logger.Error(ctx, "Ошибка при отправке сообщения пользователю", err)
		return
	}
}

func (s *Service) sendErrorMessage(ctx context.Context, msg dto.Message, text string) {
	message := bot_dto.Message{
		Chat: msg.ChatID, Text: text,
	}
	err := s.bot.SendMessage(message)
	if err != nil {
		logger.Error(ctx, "Ошибка при отправке сообщения пользователю", err)
		return
	}
}
