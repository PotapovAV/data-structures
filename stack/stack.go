package stack

type Data interface{}

type Stack struct {
	stack  []Data
	topIdx int
	size   int
}

func New(size int) *Stack {
	return &Stack{
		stack:  make([]Data, size),
		topIdx: -1,
		size:   size,
	}
}

func (s *Stack) Push(x Data) Data {
	if s.IsFull() {
		return nil
	}
	return s.putOnTop(x)
}

func (s *Stack) IsFull() bool {
	if s.topIdx == s.size-1 {
		return true
	}
	return false
}

func (s *Stack) putOnTop(x Data) Data {
	s.topIdx++
	s.stack[s.topIdx] = x
	return x
}

func (s *Stack) Pop() Data {
	if s.IsEmpty() {
		return nil
	}
	return s.getFromTop()
}

func (s *Stack) IsEmpty() bool {
	if s.topIdx == -1 {
		return true
	}
	return false
}

func (s *Stack) getFromTop() Data {
	s.topIdx--
	return s.stack[s.topIdx+1]
}
