package timer

import (
	"github.com/golang/mock/gomock"
	mocks "github.com/sosalejandro/timer/domain/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"testing"
)

type TimerManagerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl          *gomock.Controller
	mockTimer     *mocks.MockBaseTimer
	mockCallbacks *mocks.MockExecutable

	timerManager *TimerManager
}

func TestTimerManagerSuite(t *testing.T) {
	suite.Run(t, new(TimerManagerSuite))
}
func (s *TimerManagerSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())
	s.mockTimer = mocks.NewMockBaseTimer(s.ctrl)
	s.mockCallbacks = mocks.NewMockExecutable(s.ctrl)

	s.timerManager = NewTimerManager(s.mockTimer)
}

func (s *TimerManagerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *TimerManagerSuite) TestTimerManager_StartTimer() {
	s.mockTimer.EXPECT().Start().Times(1)
	s.mockCallbacks.EXPECT().ExecuteAtStart().Times(1)

	err := s.timerManager.StartTimer(s.mockCallbacks)
	s.NoError(err)
}

func (s *TimerManagerSuite) TestTimerManager_StartTimer_WithoutCallbacks() {
	s.mockTimer.EXPECT().Start().Times(1)
	s.mockCallbacks.EXPECT().ExecuteAtStart().Times(0)

	err := s.timerManager.StartTimer(nil)
	s.NoError(err)
}

func (s *TimerManagerSuite) TestTimerManager_StartTimer_InvalidTimer() {
	s.timerManager.timer = nil

	err := s.timerManager.StartTimer(s.mockCallbacks)
	s.Error(err)
}

func (s *TimerManagerSuite) TestTimerManager_StopTimer() {
	s.mockTimer.EXPECT().Stop().Times(1)
	s.mockCallbacks.EXPECT().ExecuteAtStop().Times(1)

	err := s.timerManager.StopTimer(s.mockCallbacks)
	s.NoError(err)
}

func (s *TimerManagerSuite) TestTimerManager_StopTimer_WithoutCallbacks() {
	s.mockTimer.EXPECT().Stop().Times(1)
	s.mockCallbacks.EXPECT().ExecuteAtStop().Times(0)

	err := s.timerManager.StopTimer(nil)
	s.NoError(err)
}

func (s *TimerManagerSuite) TestTimerManager_StopTimer_InvalidTimer() {
	s.timerManager.timer = nil

	err := s.timerManager.StopTimer(s.mockCallbacks)
	s.Error(err)
}

func (s *TimerManagerSuite) TestTimerManager_ResetTimer() {
	s.mockTimer.EXPECT().Reset().Times(1)
	s.mockCallbacks.EXPECT().ExecuteAtReset().Times(1)

	err := s.timerManager.ResetTimer(s.mockCallbacks)
	s.NoError(err)
}

func (s *TimerManagerSuite) TestTimerManager_ResetTimer_WithoutCallbacks() {
	s.mockTimer.EXPECT().Reset().Times(1)
	s.mockCallbacks.EXPECT().ExecuteAtReset().Times(0)

	err := s.timerManager.ResetTimer(nil)
	s.NoError(err)
}

func (s *TimerManagerSuite) TestTimerManager_ResetTimer_InvalidTimer() {
	s.timerManager.timer = nil

	err := s.timerManager.ResetTimer(s.mockCallbacks)
	s.Error(err)
}

func (s *TimerManagerSuite) TestTimerManager_IsTimerBlocked() {
	s.mockTimer.EXPECT().Blocked().Return(true).Times(1)

	blocked, err := s.timerManager.IsTimerBlocked()
	s.NoError(err)
	s.True(blocked)
}

func (s *TimerManagerSuite) TestTimerManager_IsTimerBlocked_InvalidTimer() {
	s.timerManager.timer = nil

	_, err := s.timerManager.IsTimerBlocked()
	s.Error(err)
}

func (s *TimerManagerSuite) TestNewTimerManager() {
	timerManager := NewTimerManager(s.mockTimer)
	s.NotNil(timerManager)
}

func (s *TimerManagerSuite) TestNewTimerManager_InvalidTimer() {
	timerManager := NewTimerManager(nil)
	s.NotNil(timerManager)
}
