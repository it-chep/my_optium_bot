package dao

import (
	"strings"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type Patient struct {
	xo.Patient
}

func (p *Patient) ToDomain() user.Patient {
	patient := user.Patient{
		ID:              p.ID,
		TgID:            p.TgID.Int64,
		FullName:        p.FullName.String,
		Sex:             user.Sex(p.Sex.Int64),
		BirthDate:       p.BirthDate.Time,
		MetricsLink:     p.MetricsLink.String,
		LastCommunicate: p.LastCommunicate,
	}

	patient.FirstName = patient.FullName
	if split := strings.Split(patient.FullName, " "); len(split) >= 2 {
		patient.FirstName = split[1]
	}
	return patient
}
