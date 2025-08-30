package admin

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/information_post"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/marketing"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user"
)

type HandlerAggregator struct {
	Users           *user.HandlerGroup
	Marketing       *marketing.HandlerGroup
	Scenarios       *scenarios.HandlerGroup
	InformationPost *information_post.HandlerGroup
}

func NewAggregator(adminModule *admin.Module) *HandlerAggregator {
	return &HandlerAggregator{
		Users:           user.NewGroup(adminModule),
		Marketing:       marketing.NewGroup(adminModule),
		Scenarios:       scenarios.NewGroup(adminModule),
		InformationPost: information_post.NewGroup(adminModule),
	}
}
