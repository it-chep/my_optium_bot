package information

import "github.com/it-chep/my_optium_bot.git/internal/module/bot/dto"

type ThemeID int64

const (
	// Unknown неизвестная тема
	Unknown ThemeID = iota
	// RequiredTheme тема - обязательный контент
	RequiredTheme
	// MotivationTheme тема мотивации
	MotivationTheme
	// PreparingToSecondTheme тема подготовки ко второму этапу
	PreparingToSecondTheme
)

type Post struct {
	ID        int64
	PatientID string
	Text      string

	// тема
	PostsThemeID ThemeID
	OrderInTheme int64

	// медиа
	MediaTgID string
	Type      dto.ContentType
}
