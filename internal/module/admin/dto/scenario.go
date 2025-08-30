package dto

import (
	"time"
)

var ScenarioNameMap = map[int64]string{
	1:  "Инит бота",
	2:  "Метрики",
	3:  "",
	4:  "Терапия",
	5:  "Обучение",
	6:  "Рекомендации",
	7:  "",
	8:  "Информация",
	9:  "Потеряшка",
	10: "Выведение на контроль",
}

type Scenario struct {
	ID       int64
	Name     string
	IsActive bool
	Delay    time.Duration
}

type PatientScenario struct {
	ID            int64
	PatientID     int64
	ChatID        int64
	ScenarioID    int64
	Step          int64
	Answered      bool
	Sent          bool
	ScheduledTime time.Time
	Active        bool
	Repeatable    bool
	CompletedAt   time.Time
}
