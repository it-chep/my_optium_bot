package posts

import (
	"context"

	informationDal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/information/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
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

func (s *Service) GetNextPost(ctx context.Context, ps dto.PatientScenario) (information.Post, error) {
	lastSentPost, err := s.informationDal.GetLastSentPost(ctx, ps.PatientID)
	if err != nil {
		return information.Post{}, err
	}

	nextPost, err := s.informationDal.GetNextPost(ctx, ps.PatientID, lastSentPost)
	if err != nil {
		return information.Post{}, err
	}

	return nextPost, nil
}

func (s *Service) MarkPostSent(ctx context.Context, patientTgID, sentPostID int64) error {
	return s.informationDal.MarkPostSent(ctx, patientTgID, sentPostID)
}

// FinishScenarioOrContinue заканчивает сценарий если он запустился 10 раз
func (s *Service) FinishScenarioOrContinue(ctx context.Context, patientTgID int64) error {
	count, err := s.informationDal.GetSentPostsCount(ctx, patientTgID)
	if err != nil {
		return err
	}

	if count == 10 {
		err = s.informationDal.FinishInformationScenario(ctx, patientTgID)
		return err
	}

	return nil
}
