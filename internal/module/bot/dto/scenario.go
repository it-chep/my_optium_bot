package dto

import (
	"time"

	"github.com/samber/lo"
)

type ScenarioName string

const (
	MetricsStart ScenarioName = "metrics-start"
	MetricsRetry ScenarioName = "metrics-retry"
)

type Scenario struct {
	ID            int64
	ScheduledTime time.Time
	Name          ScenarioName

	Steps Steps
}

func (s Scenario) StepByOrder(order int) (Step, bool) {
	return lo.Find(s.Steps, func(s Step) bool {
		return s.Order == order
	})
}

type Steps []Step

type Step struct {
	ID         int64
	ScenarioID int64
	Order      int
	Text       string
	IsFinal    bool
	NextStep   *int
	NextDelay  *time.Duration

	Buttons StepButtons
}

type StepButtons []StepButton

type StepButton struct {
	Text          string
	NextStepOrder int
}

type ContentType int8

const (
	// Unknown Текст или ничего
	Unknown ContentType = iota
	// Photo фотография
	Photo
	// Video Видео
	Video
	// VideoNote Кружок
	VideoNote
	// Audio Голосовое сообщение
	Audio
	// Document документ
	Document
)
