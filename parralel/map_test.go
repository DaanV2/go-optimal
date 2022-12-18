package parralel

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	sizes := []int{2, 4, 16, 64, 512, 1024, 4096}

	for _, size := range sizes {
		t.Run("Testing with size "+strconv.Itoa(size), func(t *testing.T) {
			Map_Tests(t, size)
		})
	}
}

func Map_Tests(t *testing.T, itemSize int) {
	type TestStruct struct {
		Name  string
		Age   int
		Index int
	}

	t.Run("Ensuring that all items are set in the slice", func(t *testing.T) {
		data := make([]int, itemSize)

		items, errs := Map(data, func(index int, item int, items []int) (TestStruct, error) {
			result := TestStruct{Name: "Test", Age: 10, Index: index}
			return result, nil
		})

		assert.Equal(t, 0, len(errs))

		for index, item := range items {
			assert.Equal(t, "Test", item.Name)
			assert.Equal(t, 10, item.Age)
			assert.Equal(t, index, item.Index)
		}
	})
}
