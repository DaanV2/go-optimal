package parallel

import (
	"sync"

	"github.com/daanv2/go-optimal"
)

// ForEach will execute the callback function for each item in the slice
func ForEach[T any](data []T, callbackFn func(index int, item T, items []T) error) []SliceError {
	errors := errorCollection{}
	max := len(data)
	targetSize := optimal.SliceChunkSize[T](max)
	wg := sync.WaitGroup{}

	for index := 0; index < max; index += targetSize {
		wg.Add(1)

		last := index + targetSize
		if last > max {
			last = max
		}
		go forEachItem(index, data[index:last], data, callbackFn, &wg, &errors)
	}

	wg.Wait()
	return errors.Get()
}

func forEachItem[T any](start int, section []T, items []T, callbackFn func(index int, item T, items []T) error, wg *sync.WaitGroup, errors *errorCollection) {
	defer wg.Done()

	for j, item := range section {
		if err := callbackFn(start+j, item, items); err != nil {
			errors.Add(err, start+j)
		}
	}
}
