package parallel

import (
	"errors"
	"sync"
)

type ErrorWithIndex struct {
	Index int
	Err   error
}

func (e ErrorWithIndex) Error() string {
	return e.Err.Error()
}

func (e ErrorWithIndex) Unwrap() error {
	return e.Err
}

// newErrorWithIndex wraps an error with its corresponding index if not nil.
func newErrorWithIndex(err error, index int) error {
	if err == nil {
		return nil
	}

	return ErrorWithIndex{
		Index: index,
		Err:   err,
	}
}

type errorCollection struct {
	errors error
	lock   sync.Mutex
}

// Append adds an error to the collection if it is non-nil.
func (e *errorCollection) Append(err error) {
	if err == nil {
		return
	}

	e.lock.Lock()
	defer e.lock.Unlock()

	e.errors = errors.Join(e.errors, err)
}

func (e *errorCollection) Get() error {
	return e.errors
}
