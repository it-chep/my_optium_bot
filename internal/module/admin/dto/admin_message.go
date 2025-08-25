package dto

type AdminType int8

const (
	// Admin сообщение админам из таблицы admin_messages
	Admin AdminType = iota
	// Doctor сообщение докторам из таблицы doctor_messages
	Doctor
)

type AdminMessage struct {
	ID           int64
	NextStep     int64
	Message      string
	Type         AdminType
	ScenarioName string
}
