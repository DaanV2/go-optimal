package optimal

// SliceSize returns the size
func SliceSize[T any]() int64 {
	itemSize := ObjectSize[T]()
	targetSize := GetTargetSize()

	return targetSize / itemSize
}

// MakeSlice returns a slice with the capacity set to the ideal size
func MakeSlice[T any]() []T {
	return make([]T, 0, SliceSize[T]())
}
