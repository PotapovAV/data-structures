package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyQueueDequeue(t *testing.T) {
	queue := New(2)
	assert.Equal(t, nil, queue.Dequeue(), "nil should be returned by queue.Dequeue()")
}

func TestEnqueue(t *testing.T) {
	queue := New(2)
	assert.Equal(t, 1, queue.Enqueue(1), "1 should be returned by queue.Enqueue(1)")
	assert.Equal(t, 2, queue.Enqueue(2), "2 should be returned by queue.Enqueue(2)")
}

func TestFullQueueEnqueue(t *testing.T) {
	queue := New(2)
	queue.Enqueue(1)
	queue.Enqueue(2)
	assert.Equal(t, nil, queue.Enqueue(3), "nil should be returned by queue.Enqueue(3)")
}

func TestDequeue(t *testing.T) {
	queue := New(2)
	queue.Enqueue(1)
	queue.Enqueue(2)
	assert.Equal(t, 1, queue.Dequeue(), "1 should be returned by queue.Dequeue()")
	assert.Equal(t, 2, queue.Dequeue(), "2 should be returned by queue.Dequeue()")
}
