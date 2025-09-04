package scenarios

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/create_admin_message"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/delete_admin_message"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/edit_scenario_delay"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/edit_step_text"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/get_admin_messages"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/get_scenario_steps"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/get_scenarios"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/scenarios/get_steps"
)

type HandlerGroup struct {
	GetAdminMessages   *get_admin_messages.Handler
	CreateAdminMessage *create_admin_message.Handler
	DeleteAdminMessage *delete_admin_message.Handler

	GetScenarios      *get_scenarios.Handler
	EditScenarioDelay *edit_scenario_delay.Handler
	GetScenarioSteps  *get_scenario_steps.Handler

	GetSteps     *get_steps.Handler
	EditStepText *edit_step_text.Handler
}

func NewGroup(adminModule *admin.Module) *HandlerGroup {
	return &HandlerGroup{
		GetAdminMessages:   get_admin_messages.NewHandler(adminModule),
		CreateAdminMessage: create_admin_message.NewHandler(adminModule),
		DeleteAdminMessage: delete_admin_message.NewHandler(adminModule),

		GetScenarios:      get_scenarios.NewHandler(adminModule),
		EditScenarioDelay: edit_scenario_delay.NewHandler(adminModule),
		GetScenarioSteps:  get_scenario_steps.NewHandler(adminModule),

		GetSteps:     get_steps.NewHandler(adminModule),
		EditStepText: edit_step_text.NewHandler(adminModule),
	}
}
