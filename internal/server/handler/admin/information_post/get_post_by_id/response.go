package get_post_by_id

type Post struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	ThemeName     string `json:"theme_name"`
	ThemeID       int64  `json:"theme_id"`
	Order         int    `json:"order"`
	Message       string `json:"message"`
	MediaID       string `json:"media_id"`
	ContentTypeID int64  `json:"content_type_id"`
}

type Response struct {
	Post Post `json:"post"`
}
