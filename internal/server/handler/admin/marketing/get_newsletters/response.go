package get_newsletters

type NewsLetter struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	StatusID   int8   `json:"status_id"`
	StatusName string `json:"status_name"`

	UsersCount int64 `json:"users_count"`
}

type Response struct {
	Newsletters []NewsLetter `json:"newsletters"`
}
