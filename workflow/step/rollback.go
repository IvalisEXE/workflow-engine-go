package step

import "context"

func rollback(
	executed []*Step,
) {

	for i := len(executed) - 1; i >= 0; i-- {

		step := executed[i]

		if step.compensation == nil {
			continue
		}

		_ = step.compensation.Execute(context.Background())

	}

}
