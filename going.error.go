package going

import (
	"errors"
)

// Errors going errors
type Errors struct {
	AlreadyInitialized error
	NotInitialized     error
	AlreadyRunning     error
	NotRunning         error
}

var (
	// Errs  errors
	Errs Errors = Errors{
		AlreadyInitialized: errors.New("Services is already initialized"),
		NotInitialized:     errors.New("Services is not initialized"),
		AlreadyRunning:     errors.New("Services is already running"),
		NotRunning:         errors.New("Services is not running"),
	}
)
