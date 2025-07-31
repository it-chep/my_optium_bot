package dto

type ScenarioName string

const (
	MetricsStart ScenarioName = "metrics-start"
	MetricsRetry ScenarioName = "metrics-retry"
)

type Scenario struct {
	ID   int64
	Name ScenarioName

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
