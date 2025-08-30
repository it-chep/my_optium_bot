package create_post_theme

type Request struct {
	Name       string `json:"name"`
	IsRequired bool   `json:"is_required"`
}
