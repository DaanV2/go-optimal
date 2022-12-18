package parralel

import (
	"fmt"
	"sync"
)

type SliceError struct {
	base error
	// The index of the element that caused the
	index int
}

func (e *SliceError) Error() string {
	return fmt.Sprintf("Error at index: %v", e.index) + ". " + e.base.Error()
}

type errorCollection struct {
	errors []SliceError
	lock   sync.Mutex
}

// Add adds an error to the collection.
func (e *errorCollection) Add(err error, index int) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.errors = append(e.errors, SliceError{base: err, index: index})
}

func (e *errorCollection) Get() []SliceError {
	return e.errors
}
