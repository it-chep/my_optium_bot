package dao

import "github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"

type User struct {
	IsDoctor   bool   `db:"is_doctor" json:"is_doctor"`
	ID         int64  `db:"id" json:"id"`
	ScenarioID *int64 `db:"scenario_id" json:"scenario_id"`
	StepOrder  *int64 `db:"step_order" json:"step_order"`
}

func (u *User) ToDomain() user.User {
	usr := user.User{
		IsDoctor: u.IsDoctor,
		ID:       u.ID,
	}
	if u.ScenarioID != nil && u.StepOrder != nil {
		usr.StepStat = &user.StepStat{
			ScenarioID: *u.ScenarioID,
			StepOrder:  *u.StepOrder,
		}
	}
	return usr
}
