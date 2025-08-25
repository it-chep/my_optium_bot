package education

import "github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"

type Post struct {
	ScenarioID int64
	StepID     int64
	MediaTgID  string
	Type       dto.ContentType
}
