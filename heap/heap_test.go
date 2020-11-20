package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testData struct {
	key  int
	data string
}

func (h testData) LessThan(otherItem interface{}) bool {
	switch otherItem.(type) {
	case testData:
		if h.key < otherItem.(testData).key {
			return true
		}
	}
	return false
}

func (h testData) SetKey(key interface{}) interface{} {
	switch key.(type) {
	case int:
		h.key = key.(int)
	}
	return h
}

func buildData(key int, data string) testData {
	return testData{
		key:  key,
		data: data,
	}
}

func TestEmptyHeap(t *testing.T) {
	heap := New(2)
	assert.Equal(t, 0, heap.HeapSize(), "0 should be returned by heap.HeapSize()")
	assert.Equal(t, nil, heap.ShowMax(), "nil should be returned by heap.ShowMax()")
	assert.Equal(t, nil, heap.ExtractMax(), "nil should be returned by heap.ExtractMax()")

	heap.Add(buildData(1, "one"))
	heap.Add(buildData(2, "two"))
	heap.Add(buildData(3, "three"))
	heap.Add(buildData(4, "four"))
	heap.Add(buildData(5, "five"))

	assert.Equal(t, 5, heap.ShowMax().(testData).key, "5 should be returned by heap.ShowMax()")
	assert.Equal(t, 5, heap.ExtractMax().(testData).key, "5 should be returned by heap.ExtractMax()")

	assert.Equal(t, 4, heap.ShowMax().(testData).key, "4 should be returned by heap.ShowMax()")
	assert.Equal(t, 4, heap.ExtractMax().(testData).key, "4 should be returned by heap.ExtractMax()")

	heap.ExtractMax()
	heap.ExtractMax()
	heap.ExtractMax()
	assert.Equal(t, nil, heap.ExtractMax(), "nil should be returned by heap.ExtractMax()")

	heap.Add(buildData(5, "five"))
	assert.Equal(t, 5, heap.ExtractMax().(testData).key, "5 should be returned by heap.ExtractMax()")
}

func TestBuildHeap(t *testing.T) {
	input := make([]Data, 6)
	for i := 0; i < 6; i++ {
		input[i] = buildData(i+1, "data")
	}
	heap := BuildHeap(input, 2)
	assert.Equal(t, 6, heap.ShowMax().(testData).key, "6 should be returned by heap.ShowMax()")
	assert.Equal(t, 6, heap.ExtractMax().(testData).key, "6 should be returned by heap.ExtractMax()")

	heap.ChangeKey(10, 0)
	assert.Equal(t, 10, heap.ExtractMax().(testData).key, "10 should be returned by heap.ShowMax()")
	assert.Equal(t, 4, heap.HeapSize(), "4 should be returned by heap.HeapSize()")
}
