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

type InformationPostListView struct {
	ID   int64
	Name string

	PostThemeName   string
	OrderInTheme    int64
	ThemeIsRequired bool
}

type PostTheme struct {
	ID              int64
	Name            string
	ThemeIsRequired bool
}
