package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	educationDto "github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/education"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type Content struct {
	xo.Content
}

func (c *Content) ToDomain() educationDto.Post {
	return educationDto.Post{
		ScenarioID: int64(c.ScenarioID),
		StepID:     int64(c.StepID),
		MediaTgID:  c.MediaTgID,
		Type:       dto.ContentType(c.ContentTypeID),
	}
}
