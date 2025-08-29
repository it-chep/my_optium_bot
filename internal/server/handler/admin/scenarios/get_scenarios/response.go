package get_scenarios

import "time"

type Scenario struct {
	ID    int64         `json:"id"`
	Name  string        `json:"name"`
	Delay time.Duration `json:"delay"`
}

type Response struct {
	Scenarios []Scenario `json:"scenarios"`
}
