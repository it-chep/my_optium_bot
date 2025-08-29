package create_post_theme

type Request struct {
	ThemeName  string `json:"theme_name"`
	IsRequired bool   `json:"is_required"`
}
