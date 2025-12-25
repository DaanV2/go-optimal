package parallel

import "sync"

// All runs all the functions in parallel and returns all the errors.
func All(calls ...func() error) error {
	errors := errorCollection{}
	wg := sync.WaitGroup{}

	for index, callFn := range calls {
		wg.Go(func() {
			errors.Append(newErrorWithIndex(callFn(),index))
		})
	}

	wg.Wait()
	return errors.Get()
}
