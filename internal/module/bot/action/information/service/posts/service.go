package posts

import (
	"context"

	informationDal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/information/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/information"
)

type Service struct {
	informationDal *informationDal.Repository
}

func NewService(informationDal *informationDal.Repository) *Service {
	return &Service{
		informationDal: informationDal,
	}
}

//func (s *Service) GetNextPost(ctx context.Context, ps dto.PatientScenario) (information.Post, error) {
//	lastSentPost, err := s.informationDal.GetLastSentPost(ctx, ps.PatientID)
//	if err != nil {
//		return information.Post{}, err
//	}
//
//	nextPost, err := s.informationDal.GetNextPost(ctx, ps.PatientID, lastSentPost)
//	if err != nil {
//		return information.Post{}, err
//	}
//
//	return nextPost, nil
//}

func (s *Service) MarkPostSent(ctx context.Context, patientTgID, sentPostID int64) error {
	return s.informationDal.MarkPostSent(ctx, patientTgID, sentPostID)
}

// FinishScenarioOrContinue заканчивает сценарий если он запустился 10 раз
func (s *Service) FinishScenarioOrContinue(ctx context.Context, patientTgID int64) error {
	count, err := s.informationDal.GetRepetitionsCount(ctx, patientTgID, 8)
	if err != nil {
		return err
	}

	if count == 14 {
		err = s.informationDal.FinishInformationScenario(ctx, patientTgID)
		return err
	}

	return nil
}

// GetRequiredPost получение поста по теме обязательный контент
func (s *Service) GetRequiredPost(ctx context.Context, patientTgID int64) (information.Post, error) {
	return s.informationDal.GetNextPostByTheme(ctx, patientTgID, information.RequiredTheme)
}

// GetMotivationPost получение поста по теме мотивация
func (s *Service) GetMotivationPost(ctx context.Context, patientTgID int64) (information.Post, error) {
	return s.informationDal.GetNextPostByTheme(ctx, patientTgID, information.MotivationTheme)
}

// GetSecondPartPost получение поста по теме следующего этапа
func (s *Service) GetSecondPartPost(ctx context.Context, patientTgID int64) (information.Post, error) {
	return s.informationDal.GetNextPostByTheme(ctx, patientTgID, information.PreparingToSecondTheme)
}

// GetAdditionalPost получение поста по дополнительной теме
func (s *Service) GetAdditionalPost(ctx context.Context, patientTgID int64) (information.Post, error) {
	return s.informationDal.GetAdditionalPost(ctx, patientTgID)
}

// GetInformationScenRepetitionsNumber получает количество выполненных сценариев на пользователе
func (s *Service) GetInformationScenRepetitionsNumber(ctx context.Context, patientTgID int64) (int64, error) {
	return s.informationDal.GetRepetitionsCount(ctx, patientTgID, 8)
}

func (s *Service) UpdateCounter(ctx context.Context, patientTgID, count int64) error {
	return s.informationDal.UpdateRepetitionsCount(ctx, patientTgID, 8, count)
}
