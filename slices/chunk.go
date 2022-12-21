package slices

import "github.com/daanv2/optimal/optimal"

// Chunk will chunk a slice into smaller slices based upon the optimal size
func Chunk[T any](items []T, callback func(items []T) error) error {
	chunk := optimal.SliceSize[T]()

	for i := 0; i < len(items); i += chunk {
		end := i + chunk

		if end > len(items) {
			end = len(items)
		}

		if err := callback(items[i:end]); err != nil {
			return err
		}
	}

	return nil
}
