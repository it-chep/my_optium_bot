package get_users

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Sex         string `json:"sex"`
	TgID        int64  `json:"tg_id"`
	MetricsLink string `json:"metrics_link"`
	Birthday    string `json:"birthday"`
}

type Response struct {
	Users []User `json:"users"`
}
