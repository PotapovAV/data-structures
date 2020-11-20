package linkedlist

import "fmt"

type Data interface{}

type List struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	data Data
	next *Node
	prev *Node
}

func New() *List {
	return &List{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func nodeCreate(x Data) *Node {
	return &Node{
		data: x,
		next: nil,
		prev: nil,
	}
}

func (l *List) Head() *Node {
	return l.head
}

func (l *List) Next(node *Node) *Node {
	if node == nil {
		return nil
	}
	return node.next
}

func (l *List) Data(node *Node) Data {
	if node == nil {
		return nil
	}
	return node.data
}

func (l *List) Update(node *Node, data Data) {
	node.data = data
}

func (l *List) Insert(x Data) {
	node := nodeCreate(x)
	if l.size == 0 {
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
	}
	l.head = node
	l.size++
}

func (l *List) Search(x Data) *Node {
	if l.size == 0 {
		return nil
	}
	unitPtr := l.head
	for unitPtr != nil {
		if unitPtr.data == x {
			return unitPtr
		}
		unitPtr = unitPtr.next
	}
	return nil
}

func (l *List) Remove(unitPtr *Node) Data {
	if (l.size == 0) || (unitPtr == nil) {
		return nil
	}
	if l.size == 1 {
		l.tail = nil
		l.head = nil
	} else {
		switch unitPtr {
		case l.head:
			l.head = unitPtr.next
			l.head.prev = nil
		case l.tail:
			l.tail = unitPtr.prev
			l.tail.next = nil
		default:
			unitPtr.prev.next = unitPtr.next
			unitPtr.next.prev = unitPtr.prev
		}
	}
	unitPtr.next = nil
	unitPtr.prev = nil
	l.size--
	return unitPtr.data
}

func (l *List) InsertFirst(x Data) {
	l.Insert(x)
}

func (l *List) InsertLast(x Data) {
	if l.size == 0 {
		l.Insert(x)
	} else {
		node := nodeCreate(x)
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
		l.size++
	}
}

func (l *List) GetFirst() Data {
	return l.Remove(l.head)
}

func (l *List) GetLast() Data {
	return l.Remove(l.tail)
}

func (l *List) Size() int {
	return l.size
}

func (l *List) PrintAll() {
	fmt.Print("head -> ")
	nextNodePtr := l.head
	for nextNodePtr != nil {
		fmt.Print(nextNodePtr.data, " ")
		nextNodePtr = nextNodePtr.next
	}
	fmt.Println("<- tail")
}
