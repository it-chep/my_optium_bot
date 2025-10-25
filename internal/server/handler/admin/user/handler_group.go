package user

import (
	"github.com/it-chep/my_optium_bot.git/internal/module/admin"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/add_post_to_patient"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/add_user_to_list"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/auth"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/delete_post_from_patient"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/delete_user"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/delete_user_from_list"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/get_user_by_id"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/get_users"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/move_2_step"
	"github.com/it-chep/my_optium_bot.git/internal/server/handler/admin/user/update_shedule_time"
)

type HandlerGroup struct {
	Auth *auth.Handler

	GetUsers    *get_users.Handler
	GetUserByID *get_user_by_id.Handler

	AddUserToList    *add_user_to_list.Handler
	AddPostToPatient *add_post_to_patient.Handler

	DeletePostFromPatient *delete_post_from_patient.Handler
	DeleteUserFromList    *delete_user_from_list.Handler

	UpdateSheduleTime *update_shedule_time.Handler
	DeleteUser        *delete_user.Handler
	Move2Step         *move_2_step.Handler
}

func NewGroup(adminModule *admin.Module) *HandlerGroup {
	return &HandlerGroup{
		Auth: auth.NewHandler(),

		GetUsers:    get_users.NewHandler(adminModule),
		GetUserByID: get_user_by_id.NewHandler(adminModule),

		AddUserToList:    add_user_to_list.NewHandler(adminModule),
		AddPostToPatient: add_post_to_patient.NewHandler(adminModule),

		DeletePostFromPatient: delete_post_from_patient.NewHandler(adminModule),
		DeleteUserFromList:    delete_user_from_list.NewHandler(adminModule),

		UpdateSheduleTime: update_shedule_time.NewHandler(adminModule),
		DeleteUser:        delete_user.NewHandler(adminModule),
		Move2Step:         move_2_step.NewHandler(adminModule),
	}
}
