package parallel

import "sync"

// All runs all the functions in parallel and returns all the errors.
func All(calls ...func() error) []SliceError {
	errors := errorCollection{}
	var wg sync.WaitGroup
	wg.Add(len(calls))

	for index, callFn := range calls {
		go func(index int, call func() error) {
			defer wg.Done()

			if err := call(); err != nil {
				errors.Add(err, index)
			}
		}(index, callFn)
	}

	wg.Wait()
	return errors.Get()
}
