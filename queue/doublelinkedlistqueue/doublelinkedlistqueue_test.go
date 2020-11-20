package doublelinkedlistqueue

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoubleLinkedListQueue(t *testing.T) {
	queue := New()
	assert.Equal(t, nil, queue.Dequeue(), "nil should be returned by queue.Dequeue()")

	assert.Equal(t, true, queue.IsEmpty(), "true should be returned by queue.IsEmpty()")

	for i := 0; i < 100; i++ {
		queue.Enqueue(i)
	}
	assert.Equal(t, 100, queue.Size(), "100 should be returned by queue.Size()")

	for i := 0; i < 100; i++ {
		assert.Equal(t, i, queue.Dequeue(), strconv.Itoa(i)+" should be returned by queue.Dequeue()")
	}

	assert.Equal(t, true, queue.IsEmpty(), "true should be returned by queue.IsEmpty()")
}
