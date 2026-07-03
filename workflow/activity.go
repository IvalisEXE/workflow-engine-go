package workflow

import "context"

type Activity interface {
	Name() string
	Execute(ctx context.Context) error
}
