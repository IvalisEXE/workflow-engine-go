package workflow

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

	total := len(wf.Activities())

	for i, activity := range wf.Activities() {

		fmt.Printf("[%d/%d] %s\n",
			i+1,
			total,
			activity.Name(),
		)

		err := activity.Execute(ctx)
		if err != nil {
			return fmt.Errorf("activity %s failed: %w", activity.Name(), err)
		}
	}

	fmt.Println("\nWorkflow completed successfully")
	return nil
}
