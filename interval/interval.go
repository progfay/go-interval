package interval

import (
	"context"
	"time"
)

// SetInterval calls a function repeatedly with a fixed time delay between each call.
// To stopping repeatedly function calling, call returned function.
func SetInterval(function func(), interval time.Duration) (cancel func()) {
	ctx, cancel := context.WithCancel(context.Background())
	SetIntervalWithContext(ctx, function, interval)
	return cancel
}

// SetIntervalWithContext calls a function repeatedly with a fixed time delay between each call.
// When ctx is cancelled, repeatedly function calling is cancelled.
func SetIntervalWithContext(ctx context.Context, function func(), interval time.Duration) {
	go func(f func(), i time.Duration) {
		for {
			timer := time.NewTimer(i)
			select {
			case <-ctx.Done():
				return
			case <-timer.C:
				f()
			}
		}
	}(function, interval)
}
