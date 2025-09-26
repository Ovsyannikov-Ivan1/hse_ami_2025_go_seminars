package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVector(t *testing.T) {
	v := New[int]()
	assert.Equal(t, 0, v.Size(), "New vector should have size 0")
	assert.True(t, v.Empty(), "New vector should be empty")
}

func TestNewVectorWithOptions(t *testing.T) {
	t.Run("Empty vector", func(t *testing.T) {
		v := New[int]()
		assert.Equal(t, 0, v.Size())
		assert.True(t, v.Empty())
	})

	t.Run("With capacity", func(t *testing.T) {
		v := New[int](WithCapacity[int](10))
		assert.Equal(t, 0, v.Size())
		assert.GreaterOrEqual(t, v.Capacity(), 10)
	})

	t.Run("With values", func(t *testing.T) {
		v := New[int](WithValues(1, 2, 3, 4, 5))
		assert.Equal(t, 5, v.Size())

		val, err := v.At(0)
		assert.NoError(t, err)
		assert.Equal(t, 1, val)

		val, err = v.At(4)
		assert.NoError(t, err)
		assert.Equal(t, 5, val)
	})

	t.Run("With size and default value", func(t *testing.T) {
		v := New[int](WithSize(5, 42))
		assert.Equal(t, 5, v.Size())

		for i := 0; i < 5; i++ {
			val, err := v.At(i)
			assert.NoError(t, err)
			assert.Equal(t, 42, val)
		}
	})

	t.Run("With fill", func(t *testing.T) {
		v := New[string](WithFill(3, "hello"))
		assert.Equal(t, 3, v.Size())

		for i := 0; i < 3; i++ {
			val, err := v.At(i)
			assert.NoError(t, err)
			assert.Equal(t, "hello", val)
		}
	})

	t.Run("From slice", func(t *testing.T) {
		slice := []float64{1.1, 2.2, 3.3}
		v := New[float64](FromSlice(slice))
		assert.Equal(t, 3, v.Size())

		val, err := v.At(1)
		assert.NoError(t, err)
		assert.Equal(t, 2.2, val)
	})

	t.Run("Multiple options - last wins", func(t *testing.T) {
		v := New[int](
			WithCapacity[int](10),
			WithValues(1, 2, 3),
		)
		assert.Equal(t, 3, v.Size())
		assert.Equal(t, 3, v.Capacity())
	})

	t.Run("Convenience constructors", func(t *testing.T) {
		intVec := NewInt(WithValues(1, 2, 3))
		assert.Equal(t, 3, intVec.Size())

		strVec := NewString(WithValues("a", "b", "c"))
		assert.Equal(t, 3, strVec.Size())

		floatVec := NewFloat64(WithValues(1.1, 2.2))
		assert.Equal(t, 2, floatVec.Size())
	})
}

func TestComplexTypesWithOptions(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("With struct values", func(t *testing.T) {
		people := []Person{
			{"Alice", 30},
			{"Bob", 25},
		}

		v := New[Person](FromSlice(people))
		assert.Equal(t, 2, v.Size())

		p, err := v.At(0)
		assert.NoError(t, err)
		assert.Equal(t, "Alice", p.Name)
		assert.Equal(t, 30, p.Age)
	})

	t.Run("With fill for structs", func(t *testing.T) {
		defaultPerson := Person{"Unknown", 0}
		v := New[Person](WithFill(3, defaultPerson))

		assert.Equal(t, 3, v.Size())

		for i := 0; i < 3; i++ {
			p, err := v.At(i)
			assert.NoError(t, err)
			assert.Equal(t, defaultPerson, p)
		}
	})
}

func TestEdgeCasesWithOptions(t *testing.T) {
	t.Run("Zero capacity", func(t *testing.T) {
		v := New[int](WithCapacity[int](0))
		assert.Equal(t, 0, v.Capacity())
	})

	t.Run("Negative capacity", func(t *testing.T) {
		v := New[int](WithCapacity[int](-5))
		assert.Equal(t, 0, v.Capacity())
	})

	t.Run("Empty values", func(t *testing.T) {
		v := New[int](WithValues[int]())
		assert.Equal(t, 0, v.Size())
	})

	t.Run("Zero size with default", func(t *testing.T) {
		v := New[int](WithSize(0, 42))
		assert.Equal(t, 0, v.Size())
	})

	t.Run("Negative size", func(t *testing.T) {
		v := New[int](WithSize(-3, 42))
		assert.Equal(t, 0, v.Size())
	})
}

func TestPushBack(t *testing.T) {
	v := New[int]()

	v.PushBack(1)
	v.PushBack(2)
	v.PushBack(3)

	assert.Equal(t, 3, v.Size())

	val, err := v.At(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = v.At(2)
	assert.NoError(t, err)
	assert.Equal(t, 3, val)
}

func TestPopBack(t *testing.T) {
	v := New[int](WithValues(1, 2, 3))

	err := v.PopBack()
	assert.NoError(t, err)
	assert.Equal(t, 2, v.Size())

	val, err := v.Back()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	v.Clear()
	err = v.PopBack()
	assert.Error(t, err)
}

func TestAt(t *testing.T) {
	v := New[string](WithValues("a", "b", "c"))

	val, err := v.At(1)
	assert.NoError(t, err)
	assert.Equal(t, "b", val)

	_, err = v.At(5)
	assert.Error(t, err)

	_, err = v.At(-1)
	assert.Error(t, err)
}

func TestFrontBack(t *testing.T) {
	v := New[int](WithValues(10, 20, 30))

	front, err := v.Front()
	assert.NoError(t, err)
	assert.Equal(t, 10, front)

	back, err := v.Back()
	assert.NoError(t, err)
	assert.Equal(t, 30, back)

	v.Clear()
	_, err = v.Front()
	assert.Error(t, err)

	_, err = v.Back()
	assert.Error(t, err)
}

func TestInsert(t *testing.T) {
	v := New[int](WithValues(1, 2, 4))

	err := v.Insert(2, 3)
	assert.NoError(t, err)
	assert.Equal(t, 4, v.Size())

	val, _ := v.At(2)
	assert.Equal(t, 3, val)

	val, _ = v.At(3)
	assert.Equal(t, 4, val)

	err = v.Insert(10, 5)
	assert.Error(t, err)
}

func TestErase(t *testing.T) {
	v := New[int](WithValues(1, 2, 3, 4))

	err := v.Erase(1)
	assert.NoError(t, err)
	assert.Equal(t, 3, v.Size())

	val, _ := v.At(1)
	assert.Equal(t, 3, val)

	err = v.Erase(5)
	assert.Error(t, err)
}

func TestClear(t *testing.T) {
	v := New[int](WithValues(1, 2, 3, 4, 5))

	v.Clear()

	assert.True(t, v.Empty())
	assert.Equal(t, 0, v.Size())
}

func TestReserve(t *testing.T) {
	v := New[int]()

	v.Reserve(10)
	assert.GreaterOrEqual(t, v.Capacity(), 10)
	assert.Equal(t, 0, v.Size())

	newCap := v.Capacity()
	v.Reserve(5)
	assert.Equal(t, newCap, v.Capacity())
}

func TestResize(t *testing.T) {
	v := New[int](WithValues(1, 2, 3))

	v.Resize(5, 0)
	assert.Equal(t, 5, v.Size())

	val, _ := v.At(3)
	assert.Equal(t, 0, val)

	val, _ = v.At(4)
	assert.Equal(t, 0, val)

	v.Resize(2, 0)
	assert.Equal(t, 2, v.Size())

	val, _ = v.At(1)
	assert.Equal(t, 2, val)
}

func TestSwap(t *testing.T) {
	v1 := New[int](WithValues(1, 2, 3))
	v2 := New[int](WithValues(4, 5, 6, 7))

	size1, cap1 := v1.Size(), v1.Capacity()
	size2, cap2 := v2.Size(), v2.Capacity()

	v1.Swap(v2)

	assert.Equal(t, size2, v1.Size())
	assert.Equal(t, cap2, v1.Capacity())
	assert.Equal(t, size1, v2.Size())
	assert.Equal(t, cap1, v2.Capacity())

	val, _ := v1.At(0)
	assert.Equal(t, 4, val)

	val, _ = v2.At(0)
	assert.Equal(t, 1, val)
}

func TestAssign(t *testing.T) {
	v := New[int](WithValues(1, 2, 3))

	v.Assign(4, 5, 6, 7)

	assert.Equal(t, 4, v.Size())

	val, _ := v.At(0)
	assert.Equal(t, 4, val)

	val, _ = v.At(3)
	assert.Equal(t, 7, val)
}

func TestData(t *testing.T) {
	v := New[int](WithValues(1, 2, 3))
	data := v.Data()

	assert.Equal(t, 3, len(data))
	assert.Equal(t, []int{1, 2, 3}, data)

	data[1] = 20
	val, _ := v.At(1)
	assert.Equal(t, 20, val)
}

func TestIteration(t *testing.T) {
	v := New[int](WithValues(10, 20, 30, 40, 50))
	expected := []int{10, 20, 30, 40, 50}

	index := 0
	for i := v.Begin(); i < v.End(); i++ {
		val, err := v.At(i)
		assert.NoError(t, err)
		assert.Equal(t, expected[index], val)
		index++
	}

	assert.Equal(t, len(expected), index)
}

func TestCapacityGrowth(t *testing.T) {
	v := New[int]()

	assert.Equal(t, 0, v.Capacity())

	v.PushBack(1)
	assert.Equal(t, 1, v.Capacity())

	v.PushBack(2)
	assert.Equal(t, 2, v.Capacity())

	v.PushBack(3)
	assert.Equal(t, 4, v.Capacity())
}

func TestStringRepresentation(t *testing.T) {
	v := New[int](WithValues(1, 2, 3))
	str := v.String()
	expected := "Vector[1 2 3]"

	assert.Equal(t, expected, str)
}

func TestPushBackWithPreallocatedCapacity(t *testing.T) {
	v := New[int](WithCapacity[int](100))

	// Push back should be faster with pre-allocated capacity
	for i := 0; i < 100; i++ {
		v.PushBack(i)
	}

	assert.Equal(t, 100, v.Size())
	assert.Equal(t, 100, v.Capacity())
}

// Benchmark tests with options
func BenchmarkPushBackWithPreallocation(b *testing.B) {
	b.Run("With capacity option", func(b *testing.B) {
		v := New[int](WithCapacity[int](b.N))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			v.PushBack(i)
		}
	})

	b.Run("Without preallocation", func(b *testing.B) {
		v := New[int]()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			v.PushBack(i)
		}
	})
}
