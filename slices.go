package optimal

import "runtime"

// SliceSize returns the size
func SliceSize[T any]() int {
	itemSize := ObjectSize[T]()
	targetSize := GetTargetSize()

	return int(targetSize / itemSize)
}

// SliceChunkSize returns an ideal chunk size to chunk a slice into
func SliceChunkSize[T any](size int) int {
	targetSize := SliceSize[T]()

	if targetSize > size {
		return size / runtime.GOMAXPROCS(0)
	}

	return targetSize
}

// MakeSlice returns a slice with the capacity set to the ideal size
func NewSlice[T any]() []T {
	return make([]T, 0, SliceSize[T]())
}

