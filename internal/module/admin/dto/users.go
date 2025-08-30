package dto

import "time"

type Sex int8

const (
	Man   Sex = 0
	Woman Sex = 1
)

func (s Sex) String() string {
	switch s {
	case Man:
		return "М"
	case Woman:
		return "Ж"
	}
	return "Не указан"
}

type User struct {
	ID     int64 // id в бд
	ChatID int64 // id чата с врачом
	TgID   int64 // tg_ID

	FullName    string    // Имя пользователя
	Sex         Sex       // пол пользователя
	MetricsLink string    // ссылка на метрики
	BirthDate   time.Time // дата рождения

	Lists []UserList // списки пользователя
}

// UserList список в котором состоит пользователь
type UserList struct {
	ID   int64
	Name string
}
