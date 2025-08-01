package information

import "context"

// Action Сценарий "Информация"
type Action struct {
}

func New() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context) error {
	return nil
}
