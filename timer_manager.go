package timer

import (
	"github.com/sosalejandro/timer/domain"
	"github.com/sosalejandro/timer/exceptions"
)

type TimerManager struct {
	timer domain.BaseTimer
}

func NewTimerManager(timer domain.BaseTimer) *TimerManager {
	return &TimerManager{
		timer: timer,
	}
}

func (tm *TimerManager) StartTimer() error {
	if tm.timer != nil {
		tm.timer.Start()
		return nil
	}

	return exceptions.ErrInvalidTimer
}

func (tm *TimerManager) StopTimer() error {
	if tm.timer != nil {
		tm.timer.Stop()
		return nil
	}

	return exceptions.ErrInvalidTimer
}

func (tm *TimerManager) ResetTimer() error {
	if tm.timer != nil {
		tm.timer.Reset()
		return nil
	}

	return exceptions.ErrInvalidTimer
}

func (tm *TimerManager) IsTimerBlocked() (bool, error) {
	if tm.timer != nil {
		return tm.timer.Blocked(), nil
	}

	return false, exceptions.ErrInvalidDuration
}
