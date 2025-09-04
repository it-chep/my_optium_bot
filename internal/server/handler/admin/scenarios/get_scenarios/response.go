package get_scenarios

type Scenario struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Delay string `json:"delay"`
}

type Response struct {
	Scenarios []Scenario `json:"scenarios"`
}
