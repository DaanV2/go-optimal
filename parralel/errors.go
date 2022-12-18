package parralel

import "sync"

type errorCollection struct {
	errors []error
	lock   sync.Mutex
}

func (e *errorCollection) Add(err error) {
	e.lock.Lock()
	defer e.lock.Unlock()

	e.errors = append(e.errors, err)
}

func (e *errorCollection) Get() []error {
	return e.errors
}
