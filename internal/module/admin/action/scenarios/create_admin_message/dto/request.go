package dto

import "github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"

type CreateMessageRequest struct {
	Message    string
	Type       dto.AdminType
	ScenarioID int64
	StepID     int64
}
