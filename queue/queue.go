package queue

type Data interface{}

type Queue struct {
	queue []Data
	head  int
	tail  int
	size  int
}

func New(size int) *Queue {
	return &Queue{
		queue: make([]Data, size+1),
		head:  0,
		tail:  0,
		size:  size + 1,
	}
}

func (q *Queue) Enqueue(data Data) Data {
	if q.IsFull() {
		return nil
	}
	q.queue[q.tail] = data
	q.tail = getNextIdx(q.tail, q.size)
	return data
}

func (q *Queue) IsFull() bool {
	if q.head == getNextIdx(q.tail, q.size) {
		return true
	}
	return false
}

func getNextIdx(currentIdx, size int) int {
	if currentIdx == size-1 {
		return 0
	}
	return currentIdx + 1
}

func (q *Queue) Dequeue() Data {
	if q.IsEmpty() {
		return nil
	}
	returnValue := q.queue[q.head]
	q.head = getNextIdx(q.head, q.size)
	return returnValue
}

func (q *Queue) IsEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}
