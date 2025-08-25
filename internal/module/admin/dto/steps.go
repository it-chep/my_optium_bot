package dto

import (
	"time"
)

type Step struct {
	ID           int64
	ScenarioID   int64
	ScenarioName string
	StepOrder    int64
	Content      string
	IsFinal      bool
	NextDelay    time.Duration
	NextStep     int64
}
