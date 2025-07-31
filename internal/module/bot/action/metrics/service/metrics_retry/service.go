package metrics_retry

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
	return nil
}
