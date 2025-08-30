package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type User struct {
	xo.Patient
	//Lists []xo.UserList `pgxload:"prefix=lists."`
}

func (usr *User) ToDomain() dto.User {

	//lists := make([]dto.UserList, 0, len(usr.Lists))
	//for _, list := range usr.Lists {
	//	lists = append(lists, dto.UserList{
	//		ID:   list.ID,
	//		Name: list.Name,
	//	})
	//}

	return dto.User{
		ID:          usr.ID,
		TgID:        usr.TgID.Int64,
		FullName:    usr.FullName.String,
		Sex:         dto.Sex(usr.Sex.Int64),
		MetricsLink: usr.MetricsLink.String,
		BirthDate:   usr.BirthDate.Time,
		//Lists:       lists,
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
