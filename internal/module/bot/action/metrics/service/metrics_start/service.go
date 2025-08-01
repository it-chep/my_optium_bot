package metrics_start

import (
	"context"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto/user"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) DoAction(ctx context.Context, usr user.User, msg dto.Message) (err error) {

	switch usr.StepStat.StepOrder {
	case 0:

	case 1:
	case 2:
	case 3:

	}

	return nil
}
