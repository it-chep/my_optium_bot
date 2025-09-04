package get_scenario_steps

type Step struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	StepOrder int64  `json:"step_order"`
}

type Response struct {
	Steps []Step `json:"steps"`
}
