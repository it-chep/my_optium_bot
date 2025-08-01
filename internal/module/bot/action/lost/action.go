package lost

import "context"

// Action Сценарий "Потеряшка"
type Action struct {
}

func NewAction() *Action {
	return &Action{}
}

func (action *Action) Do(ctx context.Context) error {
	return nil
}
