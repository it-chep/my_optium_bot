package update_shedule_time

type Request struct {
	NextDelay  string `json:"scheduled_time"`
	ScenarioID int64  `json:"scenario_id"`
}
