package scenarios

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/create_admin_message"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/delete_admin_message"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/edit_scenario_delay"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/edit_step_text"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/get_admin_messages"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/get_scenarios"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/get_steps"
)

type HandlerGroup struct {
	GetAdminMessages   *get_admin_messages.Handler
	CreateAdminMessage *create_admin_message.Handler
	DeleteAdminMessage *delete_admin_message.Handler

	GetScenarios      *get_scenarios.Handler
	EditScenarioDelay *edit_scenario_delay.Handler

	GetSteps     *get_steps.Handler
	EditStepText *edit_step_text.Handler
}

func NewGroup(adminModule *admin.Module) *HandlerGroup {
	return &HandlerGroup{
		GetAdminMessages:   get_admin_messages.NewHandler(),
		CreateAdminMessage: create_admin_message.NewHandler(),
		DeleteAdminMessage: delete_admin_message.NewHandler(),

		GetScenarios:      get_scenarios.NewHandler(),
		EditScenarioDelay: edit_scenario_delay.NewHandler(),

		GetSteps:     get_steps.NewHandler(),
		EditStepText: edit_step_text.NewHandler(),
	}
}
