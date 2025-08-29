package dto

import "github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"

type Response struct {
	UserData  dto.User
	Scenarios []dto.PatientScenario
	Posts     []dto.InformationPost
	Lists     []dto.UsersList
}
