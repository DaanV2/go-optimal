package env

import "os"

type environment[T any] struct {
	convert func(value string) (T, error)
}

func (e environment[T]) Lookup(key string, fallback T) T {
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

func (e environment[T]) MustLookup(key string) T {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("missing environment variable: " + key)
	}

	result, err := e.convert(value)
	if err != nil {
		panic("invalid environment variable: " + key)
	}

	return result
}

func (e environment[T]) Get(key string) (T, error) {
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
