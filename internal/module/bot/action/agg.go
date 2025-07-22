package action

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/action/create_doctor"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dal"
	"github.com/it-chep/my_optium_bot.git/internal/pkg/tg_bot"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Agg struct {
	CreateDoctor *create_doctor.Action
}

func NewAgg(pool *pgxpool.Pool, bot *tg_bot.Bot, common *dal.CommonDal) *Agg {
	return &Agg{
		CreateDoctor: create_doctor.NewAction(pool, bot, common),
	}
}
