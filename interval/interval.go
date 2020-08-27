package interval

import (
	"context"
	"time"
)

func SetInterval(function func(), interval time.Duration) (cancel func()) {
	ctx, cancel := context.WithCancel(context.Background())
	SetIntervalWithContext(ctx, function, interval)
	return cancel
}

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
