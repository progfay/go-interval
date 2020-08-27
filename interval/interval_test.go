package interval_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/progfay/go-interval/interval"
)

type Counter struct {
	count int64
}

func (c *Counter) Inc() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) GetCount() int64 {
	return atomic.LoadInt64(&c.count)
}

func Test_SetInterval(t *testing.T) {
	t.Run("Function called after the interval", func(t *testing.T) {
		c := Counter{}
		_ = interval.SetInterval(c.Inc, 200*time.Microsecond)
		time.Sleep(300 * time.Microsecond)
		if c.GetCount() != 1 {
			t.Error("First argument of SetInterval must be called after the interval.")
		}
	})

	t.Run("Canceling SetInterval", func(t *testing.T) {
		c := Counter{}
		cancel := interval.SetInterval(c.Inc, 100*time.Microsecond)
		cancel()
		time.Sleep(300 * time.Microsecond)
		if c.GetCount() != 0 {
			t.Error("First argument of SetInterval must be not called when cancelled.")
		}
	})
}

func Test_SetIntervalWithContext(t *testing.T) {
	t.Run("Canceling SetInterval when parent context is canceled", func(t *testing.T) {
		c := Counter{}
		ctx, cancel := context.WithCancel(context.Background())
		interval.SetIntervalWithContext(ctx, c.Inc, 100*time.Microsecond)
		cancel()
		time.Sleep(300 * time.Microsecond)
		if c.GetCount() != 0 {
			t.Error("First argument of SetInterval must be not called when parent context is cancelled.")
		}
	})
}
