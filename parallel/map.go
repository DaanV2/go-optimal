package parallel

import (
	"sync"

	"github.com/daanv2/go-optimal"
)

// Map will execute the callback function for each item in the slice and return the result
func Map[T any, U any](data []T, callbackFn func(index int, item T, items []T) (U, error)) ([]U, error) {
	errors := &errorCollection{}
	wg := &sync.WaitGroup{}
	max := len(data)
	result := make([]U, len(data))
	targetSize := optimal.SliceChunkSize[T](max)

	for index := 0; index < max; index += targetSize {
		wg.Add(1)

		last := index + targetSize
		if last > max {
			last = max
		}

		go mapItem(index, data[index:last], data, callbackFn, wg, errors, &result)
	}

	wg.Wait()
	return result, errors.Get()
}

func mapItem[T any, U any](start int, section []T, items []T, callbackFn func(index int, item T, items []T) (U, error), wg *sync.WaitGroup, errors *errorCollection, result *[]U) {
	defer wg.Done()

	for j, item := range section {
		item, err := callbackFn(start+j, item, items);
		if err != nil {
			errors.Add(err, start+j)
		} else {
			(*result)[start+j] = item
		}
	}
}
