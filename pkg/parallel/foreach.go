package parallel

import (
	"errors"
	"sync"

	"github.com/daanv2/go-optimal"
)

// ForEach will execute the callback function for each item in the slice
// and return all the collected errors.
func ForEach[T any](data []T, callbackFn func(index int, item T, items []T) error) error {
	errs := errorCollection{}
	wg := sync.WaitGroup{}
	maxLen := len(data)
	targetSize := optimal.SliceChunkSize[T](maxLen)

	for index := 0; index < maxLen; index += targetSize {
		last := min(index+targetSize, maxLen)

		wg.Go(func() {
			errs.Append(forEachItem(index, data[index:last], data, callbackFn))
		})
	}

	wg.Wait()

	return errs.Get()
}

func forEachItem[T any](start int, section, items []T, callbackFn func(index int, item T, items []T) error) error {
	var err error
	for j, item := range section {
		err = errors.Join(err, newErrorWithIndex(callbackFn(start+j, item, items), start+j))
	}

	return err
}
