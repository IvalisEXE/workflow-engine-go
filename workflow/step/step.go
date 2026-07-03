package step

import "time"

type Step struct {
	name string

	activity Activity

	compensation Activity

	retry RetryPolicy

	timeout time.Duration
}

func (s *Step) Retry(
	attempt int,
) *Step {

	s.retry.MaxAttempt = attempt

	return s
}

func (s *Step) Timeout(
	d time.Duration,
) *Step {

	s.timeout = d

	return s
}

func (s *Step) Compensate(
	a Activity,
) *Step {

	s.compensation = a

	return s
}
