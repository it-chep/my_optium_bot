package dao

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
)

type ContentTypeDao struct {
	xo.ContentType
}

type ContentTypes []ContentTypeDao

func (ct ContentTypeDao) ToDomain() dto.ContentTypeDTO {
	return dto.ContentTypeDTO{
		ID:   int64(ct.ID),
		Name: ct.Name.String,
	}
}

func (cts ContentTypes) ToDomain() []dto.ContentTypeDTO {
	domain := make([]dto.ContentTypeDTO, 0, len(cts))
	for _, ct := range cts {
		domain = append(domain, ct.ToDomain())
	}
	return domain
}
