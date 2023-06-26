//go:generate mockgen -source callback.go -destination mock/callback_mock.go -package mocks
package domain

type Executable interface {
	ExecuteAtStart()
	ExecuteAtStop()
	ExecuteAtReset()
}
