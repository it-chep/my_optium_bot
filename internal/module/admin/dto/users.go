package dto

import "time"

type User struct {
	ID     int64 // id в бд
	ChatID int64 // id чата с врачом
	TgID   int64 // tg_ID

	FullName    string    // Имя пользователя
	Sex         int64     // пол пользователя
	MetricsLink string    // ссылка на метрики
	BirthDate   time.Time // дата рождения

	Lists []UserList // списки пользователя
}

// UserList список в котором состоит пользователь
type UserList struct {
	ID   int64
	Name string
}
