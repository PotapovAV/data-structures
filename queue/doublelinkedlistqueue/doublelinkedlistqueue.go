package doublelinkedlistqueue

import (
	"github.com/PotapovAV/data-structures/linkedlist"
)

type Data interface{}

type Queue struct {
	queue *linkedlist.List
}

func New() *Queue {
	return &Queue{
		queue: linkedlist.New(),
	}
}

func (q *Queue) Enqueue(x Data) {
	q.queue.InsertFirst(x)
}

func (q *Queue) Dequeue() Data {
	return q.queue.GetLast()
}

func (q *Queue) IsEmpty() bool {
	if q.queue.Size() == 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return q.queue.Size()
}
