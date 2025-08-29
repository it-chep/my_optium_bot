package get_steps

type Step struct {
	ID           int64  `json:"id"`
	ScenarioName string `json:"scenario_name"`
	StepOrder    int64  `json:"step_order"`
	Text         string `json:"text"`
}

type Response struct {
	Steps []Step `json:"steps"`
}
