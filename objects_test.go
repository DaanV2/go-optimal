package optimal

import "testing"

func Test_ObjectSize(t *testing.T) {
	t.Run("Int64 should be 8 bytes", func(t *testing.T) {
		if ObjectSize[int64]() != 8 {
			t.Fail()
		}
	})

	t.Run("Int32 should be 4 bytes", func(t *testing.T) {
		if ObjectSize[int32]() != 4 {
			t.Fail()
		}
	})

	t.Run("Int16 should be 2 bytes", func(t *testing.T) {
		if ObjectSize[int16]() != 2 {
			t.Fail()
		}
	})

	t.Run("*Int32 should be 8 bytes", func(t *testing.T) {
		if ObjectSize[*int32]() != 8 {
			t.Fail()
		}
	})

	type TestStruct struct {
		A int32
		B int32
		C int32
	}

	t.Run("TestStruct should be 12 bytes", func(t *testing.T) {
		if ObjectSize[TestStruct]() != 12 {
			t.Fail()
		}
	})

	t.Run("*TestStruct should be 8 bytes", func(t *testing.T) {
		if ObjectSize[*TestStruct]() != 8 {
			t.Fail()
		}
	})
}
