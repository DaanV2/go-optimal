package parralel

import (
	"sync"

	"github.com/daanv2/optimal/optimal"
)

// ForEach will execute the callback function for each item in the slice
func ForEach[T any](data []T, callbackFn func(int, T, []T) error) []error {
	targetSize := int(optimal.SliceSize[T]())
	wg := sync.WaitGroup{}
	max := len(data)
	errors := errorCollection{}

	for index := 0; index < max; index += targetSize {
		wg.Add(1)

		go func(start int, section []T, items []T) {
			defer wg.Done()

			for j, item := range section {
				if err := callbackFn(start+j, item, items); err != nil {
					errors.Add(err)
				}
			}
		}(index, data[index:index+targetSize], data)
	}

	wg.Wait()
	return errors.Get()
}
