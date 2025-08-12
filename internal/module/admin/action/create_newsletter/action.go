package create_newsletter

import "context"

type Action struct {
}

func NewAction() *Action {
	return &Action{}
}

func (a *Action) Do(ctx context.Context) error {
	return nil
}
