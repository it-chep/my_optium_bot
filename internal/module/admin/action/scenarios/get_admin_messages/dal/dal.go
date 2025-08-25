package dal

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dal/dao"
	"github.com/it-chep/my_optium_bot.git/internal/module/admin/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Dal struct {
	pool *pgxpool.Pool
}

func NewDal(pool *pgxpool.Pool) *Dal {
	return &Dal{
		pool: pool,
	}
}

func (d *Dal) GetMessages(ctx context.Context) ([]dto.AdminMessage, error) {
	adminSql := `
		select * from admin_messages
	`

	doctorsSql := `
		select * from doctor_messages
	`

	var adminMessages dao.ListAdminMessageDao
	if err := pgxscan.Select(ctx, d.pool, &adminMessages, adminSql); err != nil {
		return nil, err
	}

	var doctorMessages dao.ListDoctorMessageDao
	if err := pgxscan.Select(ctx, d.pool, &doctorMessages, doctorsSql); err != nil {
		return nil, err
	}

	dtoAdminMessages := adminMessages.ToDomain()
	dtoDoctorMessages := doctorMessages.ToDomain()

	dtoAdminMessages = append(dtoAdminMessages, dtoDoctorMessages...)

	return dtoAdminMessages, nil
}
