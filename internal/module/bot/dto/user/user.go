package user

import (
	"time"
)

type User struct {
	IsDoctor bool
	ID       int64

	StepStat *StepStat
}

type StepStat struct {
	ScenarioID int64
	StepOrder  int64
}

type Sex int8

const (
	Man   Sex = 0
	Woman Sex = 1
)

type Patient struct {
	ID              int64
	TgID            int64
	FullName        string
	Sex             Sex
	BirthDate       time.Time
	MetricsLink     string
	FirstName       string
	LastCommunicate time.Time
}

func (p Patient) IsEmpty() bool {
	return len(p.FullName) == 0
}
