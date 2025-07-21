package bot

import (
	"context"

	"github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"
)

func (b *Bot) Route(ctx context.Context, msg dto.Message) error {
	switch msg.Text {
	case "/init_bot":

	}
	return nil
}
