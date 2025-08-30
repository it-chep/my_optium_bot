package get_users_lists

type List struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	UsersCount int64  `json:"users_count"`
}

type Response struct {
	Lists []List `json:"lists"`
}
