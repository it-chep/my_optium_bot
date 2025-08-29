package dao

import (
	"database/sql"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/it-chep/my_optium_bot.git/pkg/xo"
	"github.com/samber/lo"
)

type Newsletter struct {
	xo.Newsletter
}

type Newsletters []Newsletter

func (n Newsletter) ToDomain() dto.Newsletter {
	return dto.Newsletter{
		ID:              n.ID,
		RecipientsCount: n.RecipientsCount.Int64,
		Text:            n.Text,
		UsersLists: lo.Map(n.UsersLists, func(item sql.NullInt64, index int) int64 {
			return item.Int64
		}),
		UsersIds: lo.Map(n.UsersIds, func(item sql.NullInt64, index int) int64 {
			return item.Int64
		}),
		MediaID:   n.MediaID.String,
		CreatedAt: n.CreatedAt,
		SentAt:    lo.ToPtr(n.SentAt.Time),
		Name:      n.Name.String,
		StatusID:  dto.NewslettersStatus(n.StatusID.Int64),
	}
}

func (ns Newsletters) ToDomain() []dto.Newsletter {
	domain := make([]dto.Newsletter, 0, len(ns))
	for _, n := range ns {
		domain = append(domain, n.ToDomain())
	}
	return domain
}
