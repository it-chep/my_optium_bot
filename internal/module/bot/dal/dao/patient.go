package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type Patient struct {
	xo.Patient
}

func (p *Patient) ToDomain() user.Patient {
	return user.Patient{
		ID:          p.ID,
		TgID:        p.TgID.Int64,
		FullName:    p.FullName.String,
		Sex:         user.Sex(p.Sex.Int64),
		BirthDate:   p.BirthDate.Time,
		MetricsLink: p.MetricsLink.String,
	}
}
