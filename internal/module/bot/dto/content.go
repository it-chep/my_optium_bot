package dto

type ContentType int8

const (
	// Photo фотография
	Photo ContentType = iota + 1
	// Video Видео
	Video
	// Document документ
	Document
	// VideoNote Кружок
	VideoNote
	// Audio Голосовое сообщение
	Audio
)

type Content struct {
	ScenarioID int64
	StepID     int64
	MediaTgID  string
	Type       ContentType
}
