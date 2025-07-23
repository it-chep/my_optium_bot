package user

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

type Patient struct{}
