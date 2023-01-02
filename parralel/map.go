package parralel

import (
	"sync"

	"github.com/daanv2/go-optimal/optimal"
)

// Map will execute the callback function for each item in the slice and return the result
func Map[T any, U any](data []T, callbackFn func(index int, item T, items []T) (U, error)) ([]U, []SliceError) {
	errors := errorCollection{}
	max := len(data)
	result := make([]U, len(data))
	targetSize := optimal.SliceChunkSize[T](max)
	wg := sync.WaitGroup{}

	for index := 0; index < max; index += targetSize {
		wg.Add(1)

		last := index + targetSize
		if last > max {
			last = max
		}

		go func(start int, section []T, items []T) {
			defer wg.Done()

			for j, item := range section {
				item, err := callbackFn(start+j, item, items);
				if err != nil {
					errors.Add(err, start+j)
				} else {
					result[start+j] = item
				}
			}
		}(index, data[index:last], data)
	}

	wg.Wait()
	return result, errors.Get()
}
