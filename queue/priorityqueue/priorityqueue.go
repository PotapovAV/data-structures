package priorityqueue

import (
	"github.com/PotapovAV/data-structures/heap"
)

type Data heap.Data

type Queue struct {
	queue *heap.Heap
}

func New() *Queue {
	return &Queue{
		queue: heap.New(2),
	}
}

func (q *Queue) Enqueue(x Data) {
	q.queue.Add(x)
}

func (q *Queue) Dequeue() Data {
	return q.queue.ExtractMax()
}

func (q *Queue) IsEmpty() bool {
	if q.queue.HeapSize() == 0 {
		return true
	}
	return false
}
