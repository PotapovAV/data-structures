package heap

type Data interface {
	LessThan(other interface{}) bool
	SetKey(newKey interface{}) interface{}
}

type Heap struct {
	heap     []Data
	heapSize int
	d        int
}

func New(d int) *Heap {
	return &Heap{
		heap:     []Data{},
		heapSize: 0,
		d:        d,
	}
}

func BuildHeap(input []Data, d int) *Heap {
	inputClone := append(input[:0:0], input...)
	return &Heap{
		heap:     buildHeap(inputClone, len(input), d),
		heapSize: len(input),
		d:        d,
	}
}

func buildHeap(input []Data, heapSize, d int) []Data {
	for parentIdx := heapSize / d; parentIdx >= 0; parentIdx-- {
		heapify(input, parentIdx, heapSize, d)
	}
	return input
}

func heapify(input []Data, parentIdx, heapSize, d int) {
	largest := parentIdx
	for j := 1; j <= d; j++ {
		childIdx := child(parentIdx, j, d)
		if (childIdx < heapSize) && (input[largest].LessThan(input[childIdx])) {
			largest = childIdx
		}
	}
	if largest != parentIdx {
		swap(input, parentIdx, largest)
		heapify(input, largest, heapSize, d)
	}
}

func parent(childIdx, d int) int {
	if childIdx == 0 {
		return 0
	}
	return (childIdx - 1) / d
}

func child(parentIdx, childPosition, d int) int {
	return d*parentIdx + childPosition
}

func swap(input []Data, i, j int) {
	input[i], input[j] = input[j], input[i]
}

func (h *Heap) HeapSize() int {
	return h.heapSize
}

func (h *Heap) ShowMax() Data {
	if h.heapSize == 0 {
		return nil
	}
	return h.heap[0]
}

func (h *Heap) ExtractMax() Data {
	if h.heapSize == 0 {
		return nil
	}
	retVal := h.heap[0]
	swap(h.heap, 0, h.heapSize-1)
	h.heap = h.heap[:h.heapSize-1]
	h.heapSize--
	heapify(h.heap, 0, h.heapSize, h.d)
	return retVal
}

func (h *Heap) Add(item Data) {
	h.heap = append(h.heap, item)
	h.heapSize++
	buildHeap(h.heap, h.heapSize, h.d)
}

func (h *Heap) ChangeKey(newKey interface{}, i int) {
	if i < h.heapSize {
		h.heap[i] = h.heap[i].SetKey(newKey).(Data)
		for (i > 0) && (h.heap[parent(i, h.d)].LessThan(h.heap[i])) {
			swap(h.heap, parent(i, h.d), i)
			i = parent(i, h.d)
		}
	}
}
