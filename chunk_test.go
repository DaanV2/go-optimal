package optimal_test

import (
	"strconv"
	"testing"

	"github.com/daanv2/go-optimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Chunk(t *testing.T) {
	sizes := []int{2, 4, 16, 64, 512, 1024, 4096}

	for _, size := range sizes {
		t.Run("Testing with size "+strconv.Itoa(size), func(t *testing.T) {
			Chunk_Tests(t, size)
		})
	}
}

func Chunk_Tests(t *testing.T, itemSize int) {
	type TestStruct struct {
		Name  string
		Age   int
		Index int
	}

	t.Run("Ensuring that all items are set in the slice", func(t *testing.T) {
		data := make([]TestStruct, itemSize)

		err := optimal.Chunk(data, func(items []TestStruct) error {
			for index, item := range items {
				item.Name = "Test"
				item.Age = 10
				item.Index = 137
				items[index] = item
			}

			return nil
		})

		require.NoError(t, err)

		for _, item := range data {
			assert.Equal(t, "Test", item.Name)
			assert.Equal(t, 10, item.Age)
			assert.Equal(t, 137, item.Index)
		}
	})
}
