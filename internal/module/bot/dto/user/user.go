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
