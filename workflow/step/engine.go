package step

import (
	"context"
	"fmt"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Run(ctx context.Context, wf *Workflow) error {

	fmt.Printf("Workflow : %s\n\n", wf.Name())

	total := len(wf.Steps())

	for i, step := range wf.Steps() {

		fmt.Printf("[%d/%d] %s\n",
			i+1,
			total,
			step.activity.Name(),
		)

		err := step.activity.Execute(ctx)
		if err != nil {
			rollback(wf.Steps())
			return fmt.Errorf("activity %s failed: %w", step.activity.Name(), err)
		}
	}

	fmt.Println("\nWorkflow completed successfully")
	return nil
}
