//go:generate mockgen -source timer.go -destination mock/timer_mock.go -package mocks
package domain

import (
	"github.com/sosalejandro/timer/exceptions"
	"sync"
	"time"
)

type BaseTimer interface {
	Start()
	Reset()
	Blocked() bool
	Stop()
}

type Timer struct {
	// mu is a mutex that protects the blocked field
	mu sync.Mutex
	// blocked is true if the timer has expired and is blocked
	blocked bool
	// resetCh is a channel that is used to reset the timer
	resetCh chan struct{}
	// stopCh is a channel that is used to stop the timer
	stopCh chan struct{}
	// duration is the time span duration of the timer
	duration time.Duration
}

// NewTimer creates a new timer, as long as the duration is less than 6 hours, otherwise it returns an error.
func NewTimer(duration time.Duration) (*Timer, error) {
	if duration > 6*time.Hour {
		return nil, exceptions.ErrInvalidDuration
	}

	t := &Timer{
		resetCh:  make(chan struct{}),
		stopCh:   make(chan struct{}),
		duration: duration,
	}
	return t, nil
}

func (t *Timer) Start() {
	go t.run()
}

// run is a goroutine that runs the timer and blocks when the timer expires.
// It is started by NewTimer when a new timer is created.
// It mutates the blocked field of the Timer struct.
// It is stopped by the Stop method.
func (t *Timer) run() {
	timer := time.NewTimer(t.duration)

	for {
		select {
		case <-t.resetCh:
			if !timer.Stop() {
				<-timer.C
			}
			timer.Reset(t.duration)
			t.mu.Lock()
			t.blocked = false
			t.mu.Unlock()
		case <-timer.C:
			t.mu.Lock()
			t.blocked = true
			t.mu.Unlock()
		case <-t.stopCh:
			timer.Stop()
			return
		}

	}
}

// Reset resets the timer to its original duration.
func (t *Timer) Reset() {
	t.resetCh <- struct{}{}
}

// Blocked returns true if the timer has expired and is blocked.
func (t *Timer) Blocked() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.blocked
}

// Stop stops the timer allowing to gracefully shut down the program, exiting the goroutine.
func (t *Timer) Stop() {
	defer close(t.stopCh)
	defer close(t.resetCh)
}
