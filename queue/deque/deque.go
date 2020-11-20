package deque

type Data interface{}

type Deque struct {
	deque []Data
	head  int
	tail  int
	size  int
}

func New(size int) *Deque {
	return &Deque{
		deque: make([]Data, size+1),
		head:  0,
		tail:  0,
		size:  size + 1,
	}
}

func (d *Deque) EnqueueFront(data Data) Data {
	if d.IsFull() {
		return nil
	}
	d.head = moveIdxBack(d.head, d.size)
	d.deque[d.head] = data
	return data
}

func (d *Deque) EnqueueBack(data Data) Data {
	if d.IsFull() {
		return nil
	}
	d.deque[d.tail] = data
	d.tail = moveIdxAhead(d.tail, d.size)
	return data
}

func (d *Deque) IsFull() bool {
	if d.head == moveIdxAhead(d.tail, d.size) {
		return true
	}
	return false
}

func moveIdxAhead(currentIdx, size int) int {
	if currentIdx == size-1 {
		return 0
	}
	return currentIdx + 1
}

func moveIdxBack(currentIdx, size int) int {
	if currentIdx == 0 {
		return size - 1
	}
	return currentIdx - 1
}

func (d *Deque) DequeueFront() Data {
	if d.IsEmpty() {
		return nil
	}
	returnValue := d.deque[d.head]
	d.head = moveIdxAhead(d.head, d.size)
	return returnValue
}

func (d *Deque) DequeueBack() Data {
	if d.IsEmpty() {
		return nil
	}
	d.tail = moveIdxBack(d.tail, d.size)
	return d.deque[d.tail]
}

func (d *Deque) IsEmpty() bool {
	if d.head == d.tail {
		return true
	}
	return false
}
