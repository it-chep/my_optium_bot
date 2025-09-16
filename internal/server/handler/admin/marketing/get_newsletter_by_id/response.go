package get_newsletter_by_id

type NewsLetter struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`

	StatusID   int8   `json:"status_id"`
	StatusName string `json:"status_name"`

	UsersCount int64 `json:"users_count"`

	Text        string  `json:"text"`
	UsersLists  []int64 `json:"users_lists"`
	MediaID     string  `json:"media_id"`
	ContentType int8    `json:"content_type_id"`
}

type Response struct {
	Newsletters NewsLetter `json:"newsletters"`
}
