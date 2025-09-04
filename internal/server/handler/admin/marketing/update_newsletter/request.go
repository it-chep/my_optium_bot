package update_newsletter

type Request struct {
	Name          string  `json:"name"`
	Text          string  `json:"text"`
	UsersLists    []int64 `json:"users_lists"`
	MediaID       string  `json:"media_id"`
	ContentTypeID int64   `json:"content_type_id"`
}
