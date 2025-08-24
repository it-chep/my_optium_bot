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
	ID               int64
	ScenarioID       int64
	Order            int
	Text             string
	IsFinal          bool
	NextStep         *int
	NextDelay        *time.Duration
	DelayFromPatient bool

	Buttons StepButtons
}

type StepButtons []StepButton

type StepButton struct {
	Text          string
	NextStepOrder int
}
