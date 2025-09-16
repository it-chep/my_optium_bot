package dao

import (
	"database/sql"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type User struct {
	xo.Patient
}

func (usr *User) ToDomain() dto.User {
	return dto.User{
		ID:          usr.ID,
		TgID:        usr.TgID.Int64,
		FullName:    usr.FullName.String,
		Sex:         dto.Sex(usr.Sex.Int64),
		MetricsLink: usr.MetricsLink.String,
		BirthDate:   usr.BirthDate.Time,
	}
}

type Users []User

func (usrs Users) ToDomain() []dto.User {
	users := make([]dto.User, 0, len(usrs))

	for _, usr := range usrs {
		users = append(users, usr.ToDomain())
	}

	return users
}

type PatientToNewsletter struct {
	xo.Patient
	ChatID sql.NullInt64 `db:"chat_id"`
}

func (usr *PatientToNewsletter) ToDomain() dto.User {
	return dto.User{
		ID:          usr.ID,
		TgID:        usr.TgID.Int64,
		FullName:    usr.FullName.String,
		Sex:         dto.Sex(usr.Sex.Int64),
		MetricsLink: usr.MetricsLink.String,
		BirthDate:   usr.BirthDate.Time,
		ChatID:      usr.ChatID.Int64,
	}
}

type PatientsToNewsletter []PatientToNewsletter

func (usrs PatientsToNewsletter) ToDomain() []dto.User {
	users := make([]dto.User, 0, len(usrs))

	for _, usr := range usrs {
		users = append(users, usr.ToDomain())
	}

	return users
}
