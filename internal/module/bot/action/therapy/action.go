package therapy

import "context"

// Action Сценарий "Терапия"
type Action struct {
}

func NewAction() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context) (err error) {
	return nil
}
