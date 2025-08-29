package get_information_posts

type Post struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	ThemeName string `json:"theme_name"`
	Order     int64  `json:"order"`

	IsThemeRequired bool `json:"is_theme_required"`
}

type Response struct {
	Posts []Post `json:"posts"`
}
