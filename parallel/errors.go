package parallel

import (
	"errors"
	"sync"
)

type errorCollection struct {
	errors error
	lock   sync.Mutex
}

// Add adds an error to the collection.
func (e *errorCollection) Add(err error, index int) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.errors = errors.Join(e.errors, err)
}

func (e *errorCollection) Get() error {
	return e.errors
}
