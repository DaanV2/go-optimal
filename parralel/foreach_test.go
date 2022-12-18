package parralel

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ForEach(t *testing.T) {
	sizes := []int{2, 4, 16, 64, 512, 1024, 4096}

	for _, size := range sizes {
		t.Run("Testing with size "+strconv.Itoa(size), func(t *testing.T) {
			ForEach_Tests(t, size)
		})
	}
}

func ForEach_Tests(t *testing.T, itemSize int) {
	type TestStruct struct {
		Name string
		Age  int
	}

	t.Run("Ensuring that all items are set in the slice", func(t *testing.T) {
		data := make([]TestStruct, itemSize)

		errs := ForEach(data, func(index int, item TestStruct, items []TestStruct) error {
			items[index] = TestStruct{
				Name: "Test",
				Age:  10,
			}
			return nil
		})

		assert.Equal(t, 0, len(errs))

		for _, item := range data {
			assert.Equal(t, "Test", item.Name)
			assert.Equal(t, 10, item.Age)
		}
	})
}
