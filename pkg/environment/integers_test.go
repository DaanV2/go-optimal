package env_test

import (
	"testing"

	env "github.com/daanv2/go-optimal/pkg/environment"
	"github.com/stretchr/testify/assert"
)

func Test_Integers(t *testing.T) {
	t.Setenv("TESTING_VALUE", "123")

	t.Run("Testing int64", func(t *testing.T) {
		value := env.Int64.Lookup("TESTING_VALUE", 0)
		assert.Equal(t, int64(123), value)
	})

	t.Run("Testing int32", func(t *testing.T) {
		value := env.Int32.Lookup("TESTING_VALUE", 0)
		assert.Equal(t, int32(123), value)
	})

	t.Run("Testing int16", func(t *testing.T) {
		value := env.Int16.Lookup("TESTING_VALUE", 0)
		assert.Equal(t, int16(123), value)
	})

	t.Run("Testing int8", func(t *testing.T) {
		value := env.Int8.Lookup("TESTING_VALUE", 0)
		assert.Equal(t, int8(123), value)
	})

	t.Run("Testing int", func(t *testing.T) {
		value := env.Int.Lookup("TESTING_VALUE", 0)
		assert.Equal(t, 123, value)
	})

	t.Run("Testing int64 default", func(t *testing.T) {
		value := env.Int64.Lookup("TESTING_VALUE_2", 789)
		assert.Equal(t, int64(789), value)
	})

	t.Run("Testing int32 default", func(t *testing.T) {
		value := env.Int32.Lookup("TESTING_VALUE_2", 789)
		assert.Equal(t, int32(789), value)
	})

	t.Run("Testing int16 default", func(t *testing.T) {
		value := env.Int16.Lookup("TESTING_VALUE_2", 789)
		assert.Equal(t, int16(789), value)
	})

	t.Run("Testing int8 default", func(t *testing.T) {
		value := env.Int8.Lookup("TESTING_VALUE_2", 102)
		assert.Equal(t, int8(102), value)
	})
}
