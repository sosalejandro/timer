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

func (tm *TimerManager) StartTimer(cb domain.Executable) error {
	if tm.timer == nil {
		return exceptions.ErrInvalidTimer
	}

	tm.timer.Start()

	if cb != nil {
		cb.ExecuteAtStart()
	}

	return nil
}

func (tm *TimerManager) StopTimer(cb domain.Executable) error {
	if tm.timer == nil {
		return exceptions.ErrInvalidTimer
	}

	tm.timer.Stop()
	if cb != nil {
		cb.ExecuteAtStop()
	}

	return nil
}

func (tm *TimerManager) ResetTimer(cb domain.Executable) error {
	if tm.timer == nil {
		return exceptions.ErrInvalidTimer
	}

	tm.timer.Reset()
	if cb != nil {
		cb.ExecuteAtReset()
	}

	return nil
}

func (tm *TimerManager) IsTimerBlocked() (bool, error) {
	if tm.timer != nil {
		return tm.timer.Blocked(), nil
	}

	return false, exceptions.ErrInvalidTimer
}
