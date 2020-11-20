package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type QueueData struct {
	priority int
	data     string
}

func (h QueueData) LessThan(other interface{}) bool {
	switch other.(type) {
	case QueueData:
		if h.priority < other.(QueueData).priority {
			return true
		}
	}
	return false
}

func (h QueueData) SetKey(key interface{}) interface{} {
	switch key.(type) {
	case int:
		h.priority = key.(int)
	}
	return h
}

func buildQueueData(priority int, data string) QueueData {
	return QueueData{
		priority: priority,
		data:     data,
	}
}

func TestEmptyDequeue(t *testing.T) {
	queue := New()
	assert.Equal(t, nil, queue.Dequeue(), "nil should be returned by queue.Dequeue()")
	assert.Equal(t, true, queue.IsEmpty(), "true should be returned by queue.IsEmpty()")
}

func TestPriorityQueue(t *testing.T) {
	queue := New()
	queue.Enqueue(buildQueueData(1, "first"))
	queue.Enqueue(buildQueueData(2, "second"))
	queue.Enqueue(buildQueueData(3, "third"))
	assert.Equal(t, 3, queue.Dequeue().(QueueData).priority, "3 should be returned by queue.Dequeue()")
	assert.Equal(t, 2, queue.Dequeue().(QueueData).priority, "2 should be returned by queue.Dequeue()")
	queue.Enqueue(buildQueueData(4, "fourth"))
	assert.Equal(t, 4, queue.Dequeue().(QueueData).priority, "4 should be returned by queue.Dequeue()")
	assert.Equal(t, 1, queue.Dequeue().(QueueData).priority, "1 should be returned by queue.Dequeue()")
}
