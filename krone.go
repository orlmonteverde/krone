package krone

import (
	"context"
	"time"
)

func defaultFilterFunc(t time.Time) bool { return true }

// Krone handle time events.
type Krone struct {
	duration   time.Duration
	FilterFunc func(t time.Time) bool
	ticker     *time.Ticker
	timer      *time.Timer
	ctx        context.Context
	cancel     context.CancelFunc
}

// Do execute a function when duration time.
func (k *Krone) Do(f func()) {
	k.timer = time.NewTimer(k.duration)

	select {
	case <-k.timer.C:
		f()
	case <-k.ctx.Done():
		return
	}
}

// NewWithContext returns a Krone instance with context.
func NewWithContext(ctx context.Context, d time.Duration) *Krone {
	ctx, cancel := context.WithCancel(ctx)
	return &Krone{
		duration: d,
		ctx:      ctx,
		cancel:   cancel,
	}
}

// New returns a Krone instance with duration.
func New(d time.Duration) *Krone {
	return NewWithContext(context.Background(), d)
}

// FromTimeWithContext returns a Krone instance with context.
func FromTimeWithContext(ctx context.Context, t time.Time) *Krone {
	ctx, cancel := context.WithCancel(ctx)
	return &Krone{
		duration: time.Until(t),
		ctx:      ctx,
		cancel:   cancel,
	}
}

// FromTime returns a Krone instance with time.
func FromTime(t time.Time) *Krone {
	return FromTimeWithContext(context.Background(), t)
}

// Every execute a function each time interval.
func (k *Krone) Every(f func()) {
	k.ticker = time.NewTicker(k.duration)

	if k.FilterFunc == nil {
		k.FilterFunc = defaultFilterFunc
	}

	for {
		select {
		case t := <-k.ticker.C:
			if k.FilterFunc(t) {
				f()
			}
		case <-k.ctx.Done():
			return
		}
	}
}

// Stop the internal ticker/timer.
func (k *Krone) Stop() {
	if k.ticker != nil {
		k.ticker.Stop()
	}

	if k.timer != nil {
		k.timer.Stop()
	}

	k.cancel()
}
