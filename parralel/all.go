package parralel

import "sync"

func All(calls ...func() error) []error {
	errors := errorCollection{}
	var wg sync.WaitGroup
	wg.Add(len(calls))

	for _, callFn := range calls {
		go func(call func() error) {
			defer wg.Done()

			if err := call(); err != nil {
				errors.Add(err)
			}
		}(callFn)
	}

	wg.Wait()
	return errors.Get()
}
