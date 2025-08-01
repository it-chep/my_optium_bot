package action

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/create_doctor"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/doctors/init_chat"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/invite_patient"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/metrics"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Agg struct {
	CreateDoctor  *create_doctor.Action
	InitChat      *init_chat.Action
	InvitePatient *invite_patient.Action
	Metrics       *metrics.Action
}

func NewAgg(pool *pgxpool.Pool, bot *tg_bot.Bot, common *dal.CommonDal) *Agg {
	return &Agg{
		CreateDoctor:  create_doctor.NewAction(pool, bot, common),
		InitChat:      init_chat.NewAction(pool, bot, common),
		InvitePatient: invite_patient.NewAction(pool, bot, common),
		Metrics:       metrics.NewAction(),
	}
}
