package bot_dto

import "github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"

type Message struct {
	Chat    int64
	Text    string
	Buttons dto.StepButtons
}
