package action

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/add_media"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/create_admin_chat"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/create_doctor"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/exit"
	exitDal "github.com/it-chep/my_optium_bot.git/internal/module/bot/action/commands/exit/dal"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/doctors/init_chat"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/education"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/invite_patient"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/text_handler"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Agg struct {
	CreateDoctor  *create_doctor.Action
	InitChat      *init_chat.Action
	InvitePatient *invite_patient.Action

	// Сценарии админа
	AddMedia        *add_media.Action
	Exit            *exit.Action
	CreateAdminChat *create_admin_chat.Action

	// Сценарии пациента
	TextHandler *text_handler.Action
	Education   *education.Action
}

func NewAgg(pool *pgxpool.Pool, bot *tg_bot.Bot, common *dal.CommonDal) *Agg {
	return &Agg{
		CreateDoctor:  create_doctor.NewAction(pool, bot, common),
		InitChat:      init_chat.NewAction(pool, bot, common),
		InvitePatient: invite_patient.NewAction(pool, bot, common),
		TextHandler:   text_handler.NewAction(common, bot),
		Education:     education.NewAction(pool, bot, common),

		// Сценарии админа
		AddMedia:        add_media.New(bot, common),
		Exit:            exit.NewAction(bot, exitDal.NewDal(pool)),
		CreateAdminChat: create_admin_chat.NewAction(pool, bot),
	}
}
