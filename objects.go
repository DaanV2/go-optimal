package optimal

import "unsafe"

// ObjectSize returns the size of a object.
// If the T is a pointer it will return the size of the pointer
func ObjectSize[T any]() int64 {
	var result T

	return int64(unsafe.Sizeof(result))
}
