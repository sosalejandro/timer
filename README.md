# Timer

The package contains a timer class.

The purpose of this class is to manage timeframes, if the state of the `Timer.Blocked()` resolves to `true` it means the timer duration has expired, therefore a flow of logic could be triggered.

The idea I had in mind for creating this package was call the `Reset()` everytime I perform an action and if no action has been triggered during the initial duration of the timer the `Timer.Blocked()` 

Sort of cheating an observables' implementation. 


## Usage

```go
package main

import (
    "time"
    "github.com/sosalejandro/timer"
)

func main() {
    t, err := timer.NewTimer(time.Second * 5)
	
	if err != nil {
		// handle error
    }
    
	if t.Blocked() {
		// trigger a flow...
    }
	
	// perform some actions ...
	t.Reset()

	// perform some actions ...
	t.Reset()    
    
    
    // Close its internal channels and exit the goroutine
	t.Stop()
}
```