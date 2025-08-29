package get_admin_messages

type Message struct {
	ID           int64  `json:"id"`
	ScenarioName string `json:"scenario_name"`
	Type         int8   `json:"type"`
	TypeName     string `json:"type_name"`
	Text         string `json:"text"`
}

type Response struct {
	Messages []Message `json:"messages"`
}
