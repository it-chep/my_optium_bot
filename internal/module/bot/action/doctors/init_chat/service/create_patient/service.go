package create_patient

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
)

type Storage interface {
	CreatePatient(ctx context.Context, fullName string) (int64, error)
	CreateM2MPatientDoctor(ctx context.Context, chatID, doctorTgID, patientID int64) error
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) CreatePatient(ctx context.Context, msg dto.Message) error {
	// Устанавливаем фио пациента
	// todo транзакцию ?

	var patientID int64
	patientID, err := s.storage.CreatePatient(ctx, msg.Text)
	if err != nil {
		return err
	}
	err = s.storage.CreateM2MPatientDoctor(ctx, msg.ChatID, msg.User, patientID)
	if err != nil {
		return err
	}

	return nil
}
