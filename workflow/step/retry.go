package step

import "time"

type RetryPolicy struct {
	MaxAttempt int

	Delay time.Duration
}
