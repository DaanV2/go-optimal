package parralel

import (
	"sync"

	"github.com/daanv2/optimal/optimal"
)

// Map will execute the callback function for each item in the slice and return the result
func Map[T any](data []T, callbackFn func(int, T, []T) (T, error)) ([]T, []error) {
	targetSize := int(optimal.SliceSize[T]())
	wg := sync.WaitGroup{}
	max := len(data)
	errors := errorCollection{}
	result := make([]T, len(data))

	for index := 0; index < max; index += targetSize {
		wg.Add(1)

		go func(start int, section []T, items []T) {
			defer wg.Done()

			for j, item := range section {
				item, err := callbackFn(start+j, item, items);
				if err != nil {
					errors.Add(err)
				} else {
					result[start+j] = item
				}
			}
		}(index, data[index:index+targetSize], data)
	}

	wg.Wait()
	return result, errors.Get()
}
