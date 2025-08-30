package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type AdminMessageDao struct {
	xo.AdminMessage
}

func (amd *AdminMessageDao) ToDomain() dto.AdminMessage {
	return dto.AdminMessage{
		ID:           int64(amd.ID),
		NextStep:     int64(amd.NextStep),
		Message:      amd.Message,
		Type:         dto.Admin,
		ScenarioName: dto.ScenarioNameMap[amd.ScenarioID],
	}
}

type ListAdminMessageDao []AdminMessageDao

func (lamd ListAdminMessageDao) ToDomain() []dto.AdminMessage {
	domains := make([]dto.AdminMessage, 0, len(lamd))
	for _, adminMessage := range lamd {
		domains = append(domains, adminMessage.ToDomain())
	}

	return domains
}

type DoctorMessageDao struct {
	xo.DoctorMessage
}

func (dmd *DoctorMessageDao) ToDomain() dto.AdminMessage {
	return dto.AdminMessage{
		ID:           int64(dmd.ID),
		NextStep:     int64(dmd.NextStep),
		Message:      dmd.Message,
		Type:         dto.Doctor,
		ScenarioName: dto.ScenarioNameMap[dmd.ScenarioID],
	}
}

type ListDoctorMessageDao []DoctorMessageDao

func (ldmd ListDoctorMessageDao) ToDomain() []dto.AdminMessage {
	domains := make([]dto.AdminMessage, 0, len(ldmd))
	for _, adminMessage := range ldmd {
		domains = append(domains, adminMessage.ToDomain())
	}

	return domains
}
