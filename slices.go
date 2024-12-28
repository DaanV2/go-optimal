package optimal

import (
	"runtime"

	"github.com/daanv2/go-optimal/pkg/cpu"
)

// SliceSize returns the size that a slice should be to fit in the target cache
func SliceSize[T any]() int {
	itemSize := ObjectSize[T]()
	targetSize := GetTargetSize()

	return int(targetSize / itemSize)
}

// SliceSizeFor returns the size that a slice should be to fit in the target cache
func SliceSizeFor[T any](kind cpu.CacheKind) int {
	itemSize := ObjectSize[T]()
	targetSize := GetTargetSizeFor(kind)

	return int(targetSize / itemSize)
}

// SliceChunkSize returns an ideal chunk size to chunk a slice into
func SliceChunkSize[T any](size int) int {
	targetSize := SliceSize[T]()

	if targetSize > size {
		targetSize = size / runtime.GOMAXPROCS(0)
	}

	return max(targetSize, 1)
}

// MakeSlice returns a slice with the capacity set to the ideal size
func NewSlice[T any]() []T {
	return make([]T, 0, SliceSize[T]())
}
