package create_admin_message

type Request struct {
	Message    string `json:"message"`
	Type       int8   `json:"type"`
	ScenarioID int64  `json:"scenario_id"`
	StepID     int64  `json:"step_id"`
}
