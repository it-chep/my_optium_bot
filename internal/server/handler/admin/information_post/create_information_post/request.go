package create_information_post

type Request struct {
	PostName      string `json:"post_name"`
	ThemeID       int64  `json:"theme_id"`
	Order         int64  `json:"order"`
	MediaID       string `json:"media_id"`
	ContentTypeID int64  `json:"content_type_id"`
	Message       string `json:"message"`
}
