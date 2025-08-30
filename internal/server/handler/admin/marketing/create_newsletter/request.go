package create_newsletter

type Request struct {
	Name          string  `json:"name"`
	UsersList     []int64 `json:"users_lists"`
	Text          string  `json:"text"`
	MediaID       string  `json:"media_id"`
	ContentTypeID int64   `json:"content_type_id"`
}
