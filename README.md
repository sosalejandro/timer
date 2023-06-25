# Timer

# Timer Implementation

The package `domain` contains a timer class.

The purpose of this class is to manage timeframes, if the state of the `Timer.Blocked()` resolves to `true` it means the timer duration has expired, therefore a flow of logic could be triggered.

The idea I had in mind for creating this package was call the `Reset()` everytime I perform an action and if no action has been triggered during the initial duration of the timer the `Timer.Blocked()` 

Sort of cheating an observables' implementation. 

# Timer Manager 

The package `timer` contains a timer manager class.

It is a wrapper of the `domain.Timer` class, it is in charge of starting the timer and managing its state.

Not necessarily the timer manager has to be used, it is just a wrapper to make the timer more friendly to use, allowing quick modifications for future implementations.

The Timer Manager exposes the following methods:
- `StartTimer()`: Starts the timer.
- `StopTimer()`: Stops the timer.
- `ResetTimer()`: Resets the timer.
- `IsTimerBlocked()`: Returns the state of the timer.

** Once the timer is stopped it must be started again in order to be used. So call the StopTimer() method only when you are sure you won't need the timer anymore.

## Usage

```go
package main

import (
	"time"

	"github.com/sosalejandro/timer"
	"github.com/sosalejandro/timer/domain"
)

func main() {
	t, err := domain.NewTimer(time.Second * 5)

	if err != nil {
		// handle error
	}

	tM := timer.NewTimerManager(t)

	// Start timer
	if err := tM.StartTimer(); err != nil {
		// handle error
	}

	// before performing an action first check the timer isn't blocked ...
	ok, err := tM.IsTimerBlocked()

	if err != nil {
		// handle error
	}

	if ok {
		// trigger a flow...
	}

	// perform some actions ...
	if err := tM.ResetTimer(); err != nil {
		// handle error
	}

	// perform some actions ...
	if err := tM.ResetTimer(); err != nil {
		// handle error
	}

	// Close its internal channels and exit the goroutine
	if err := tM.StopTimer(); err != nil {
		// handle error
	}
}

```