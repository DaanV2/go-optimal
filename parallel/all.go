package parallel

import "sync"

// All runs all the functions in parallel and returns all the errors.
func All(calls ...func() error) error {
	errors := &errorCollection{}
	wg := &sync.WaitGroup{}
	wg.Add(len(calls))

	for index, callFn := range calls {
		go allItem(index, callFn, wg, errors)
	}

	wg.Wait()
	return errors.Get()
}

func allItem(index int, call func() error, wg *sync.WaitGroup, errors *errorCollection) {
	defer wg.Done()

	if err := call(); err != nil {
		errors.Add(err, index)
	}
}
