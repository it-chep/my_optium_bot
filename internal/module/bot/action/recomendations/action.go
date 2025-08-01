package recomendations

import "context"

// Action Сценарий "Рекомендации"
type Action struct {
}

func NewAction() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context) error {
	return nil
}
