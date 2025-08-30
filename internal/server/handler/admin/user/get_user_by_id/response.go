package get_user_by_id

type UserData struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Sex         string `json:"sex"`
	TgID        int64  `json:"tg_id"`
	MetricsLink string `json:"metrics_link"`
	Birthday    string `json:"birthday"`
}

type List struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Post struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	IsRequiredTheme bool   `json:"is_required_theme"`
}

type Scenario struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	NextDelay string `json:"next_delay"`
}

type Response struct {
	User      UserData   `json:"user"`
	Lists     []List     `json:"lists"`
	Posts     []Post     `json:"posts"`
	Scenarios []Scenario `json:"scenarios"`
}
