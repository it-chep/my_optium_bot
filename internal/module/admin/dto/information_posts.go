package dto

type InformationPost struct {
	ID            int64
	Name          string
	PostsThemeID  int64
	OrderInTheme  int
	MediaID       string
	ContentTypeID int64
	PostText      string

	ThemeIsRequired bool
}
