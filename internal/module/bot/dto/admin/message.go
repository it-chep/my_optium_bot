package admin

type MessageAdmin struct {
	ChatIDs  []int64
	Messages []string
}

type MessageDoctor struct {
	DoctorID int64
	Messages []string
}
