package dto

import "time"

type Scenario struct {
	ID            int64
	ScheduledTime time.Time

	Steps Steps
}

type Steps []Step

type Step struct {
	ID         int64
	ScenarioID int64
	Order      int
	Text       string
	IsFinal    bool

	Buttons StepButtons
}

type StepButtons []StepButton

type StepButton struct {
	Text string
}
