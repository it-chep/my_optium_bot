package get_posts_themes

type Theme struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	IsThemeRequired bool   `json:"is_theme_required"`
}

type Response struct {
	Themes []Theme `json:"themes"`
}
