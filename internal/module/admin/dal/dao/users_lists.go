package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

// UserList список пользователей
type UserList struct {
	xo.UserList
	UsersCount int64 `db:"users_count" json:"users_count"` // users_count
}

type UserLists []UserList

func (urls UserLists) ToDomain() []dto.UsersList {
	usersLists := make([]dto.UsersList, 0, len(urls))
	for _, u := range urls {
		usersLists = append(usersLists, dto.UsersList{
			UsersCount: u.UsersCount,
			UserList: dto.UserList{
				ID:   u.ID,
				Name: u.Name,
			},
		})
	}

	return usersLists
}
