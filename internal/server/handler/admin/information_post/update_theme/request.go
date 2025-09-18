package update_theme

type Request struct {
	Name       string `json:"name"`
	IsRequired bool   `json:"is_theme_required"`
}
