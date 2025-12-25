package parallel

import (
	"errors"
	"sync"

	"github.com/daanv2/go-optimal"
)

// Map will execute the callback function for each item in the slice and return the result
func Map[T any, U any](data []T, callbackFn func(index int, item T, items []T) (U, error)) ([]U, error) {
	errs := errorCollection{}
	wg := sync.WaitGroup{}
	maxLen := len(data)
	result := make([]U, len(data))
	targetSize := optimal.SliceChunkSize[T](maxLen)

	for index := 0; index < maxLen; index += targetSize {
		last := min(index+targetSize, maxLen)
		wg.Go(func() {
			errs.Append(mapItem(index, data[index:last], data, callbackFn, &result))
		})
	}

	wg.Wait()

	return result, errs.Get()
}

func mapItem[T any, U any](start int, section, items []T, callbackFn func(index int, item T, items []T) (U, error), result *[]U) error {
	var err error

	for j, item := range section {
		item, cerr := callbackFn(start+j, item, items)
		if cerr != nil {
			err = errors.Join(err, newErrorWithIndex(cerr, start+j))
		} else {
			(*result)[start+j] = item
		}
	}

	return err
}
