package update_patient

import (
	"context"
	"time"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
)

type Storage interface {
	UpdatePatientSex(ctx context.Context, patientID int64, sex user.Sex) error
	UpdatePatientBirthDate(ctx context.Context, patientID int64, birthDate time.Time) error
	UpdatePatientMetricsLink(ctx context.Context, patientID int64, metricsLink string) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) UpdateSex(ctx context.Context, patientID int64, msg dto.Message) error {
	// Устанавливаем Пол пациента
	var sex user.Sex
	sex = user.Man
	if msg.Text == "Ж" {
		sex = user.Woman
	}

	err := s.storage.UpdatePatientSex(ctx, patientID, sex)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateBirthDate(ctx context.Context, patientID int64, msg dto.Message) error {
	// Устанавливаем дату рождения пациента
	var birthDate time.Time
	birthDate, err := time.Parse("02.01.2006", msg.Text)
	if err != nil {
		return err
	}

	err = s.storage.UpdatePatientBirthDate(ctx, patientID, birthDate)
	if err != nil {
		return err
	}

	return nil
}
