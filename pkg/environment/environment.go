package env

import "os"

// EnvironmentType is a type that can be used to get environment variables and convert them to T.
type EnvironmentType[T any] struct {
	convert func(value string) (T, error)
}

// Lookup returns the value of the EnvironmentType variable named by the key. converted to T.
// If the variable is not present or error, the fallback value is returned.
func (e EnvironmentType[T]) Lookup(key string, fallback T) T {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	result, err := e.convert(value)
	if err != nil {
		return fallback
	}

	return result
}

// MustLookup returns the value of the EnvironmentType variable named by the key. converted to T.
// If the variable is not present or error, the program panics.
func (e EnvironmentType[T]) MustLookup(key string) T {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("missing EnvironmentType variable: " + key)
	}

	result, err := e.convert(value)
	if err != nil {
		panic("invalid EnvironmentType variable: " + key)
	}

	return result
}

// LookupIf returns the value of the EnvironmentType variable named by the key. converted to T.
// If the variable is not present or error, the default value of T is returned and false is returned.
func (e EnvironmentType[T]) LookupIf(key string) (T, bool) {
	var result T
	value, ok := os.LookupEnv(key)
	if !ok {
		return result, false
	}

	result, err := e.convert(value)
	if err != nil {
		return result, false
	}

	return result, true
}

// Get returns the value of the EnvironmentType variable named by the key. converted to T.
// If the variable is not present or error, the default value of T is returned.
func (e EnvironmentType[T]) Get(key string) (T, error) {
	var (
		value T
		ok    bool
		str   string
	)

	str, ok = os.LookupEnv(key)
	if !ok {
		return value, os.ErrNotExist
	}

	return e.convert(str)
}
