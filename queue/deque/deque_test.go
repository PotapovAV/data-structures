package deque

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyDeuqeMixedDequeue(t *testing.T) {
	deque := New(2)
	assert.Equal(t, nil, deque.DequeueFront(), "nil should be returned by deque.DequeueFront()")
	assert.Equal(t, nil, deque.DequeueBack(), "nil should be returned by deque.DequeueBack()")
}

func TestMixedEnqueue(t *testing.T) {
	deque := New(4)
	assert.Equal(t, 1, deque.EnqueueFront(1), "1 should be returned by deque.EnqueueFront(1)")
	assert.Equal(t, 2, deque.EnqueueBack(2), "2 should be returned by deque.EnqueueBack(2)")
	assert.Equal(t, 3, deque.EnqueueFront(3), "3 should be returned by deque.EnqueueFront(3)")
	assert.Equal(t, 4, deque.EnqueueBack(4), "4 should be returned by deque.EnqueueBack(4)")
}

func TestMixedDequeue(t *testing.T) {
	deque := New(2)
	deque.EnqueueFront(1)
	deque.EnqueueFront(2)
	assert.Equal(t, 2, deque.DequeueFront(), "2 should be returned by deque.DequeueFront()")
	assert.Equal(t, 1, deque.DequeueFront(), "1 should be returned by deque.DequeueFront()")
	deque.EnqueueFront(1)
	deque.EnqueueFront(2)
	assert.Equal(t, 1, deque.DequeueBack(), "1 should be returned by deque.DequeueBack()")
	assert.Equal(t, 2, deque.DequeueBack(), "2 should be returned by deque.DequeueBack()")
	deque.EnqueueFront(1)
	deque.EnqueueFront(2)
	assert.Equal(t, 2, deque.DequeueFront(), "2 should be returned by deque.DequeueFront()")
	assert.Equal(t, 1, deque.DequeueBack(), "1 should be returned by deque.DequeueBack()")
}

func TestFullEnqueueMixed(t *testing.T) {
	deque := New(2)
	deque.EnqueueFront(1)
	deque.EnqueueFront(2)
	assert.Equal(t, nil, deque.EnqueueFront(3), "nil should be returned by deque.EnqueueFront(3)")
	assert.Equal(t, nil, deque.EnqueueBack(3), "nil should be returned by deque.EnqueueBack(3)")
}
