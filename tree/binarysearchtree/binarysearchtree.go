package binarysearchtree

type Data interface {
	LessThan(other interface{}) bool
	EqualsTo(other interface{}) bool
}

type Node struct {
	data                Data
	parent, left, right *Node
}

func newNode(data Data) *Node {
	return &Node{
		data:   data,
		parent: nil,
		left:   nil,
		right:  nil,
	}
}

func (n *Node) Data() Data {
	return n.data
}

func (n *Node) Successor() *Node {
	if n.right != nil {
		return minimum(n.right)
	}
	parentPtr := n.parent
	for (parentPtr != nil) && (n == parentPtr.right) {
		n = parentPtr
		parentPtr = parentPtr.parent
	}
	return parentPtr
}

func minimum(nextNodePtr *Node) *Node {
	var minNodePtr *Node = nil
	for nextNodePtr != nil {
		minNodePtr = nextNodePtr
		nextNodePtr = nextNodePtr.left
	}
	return minNodePtr
}

func (n *Node) Predecessor() *Node {
	if n.left != nil {
		return maximum(n.left)
	}
	parentPtr := n.parent
	for (parentPtr != nil) && (n == parentPtr.left) {
		n = parentPtr
		parentPtr = parentPtr.parent
	}
	return parentPtr
}

func maximum(nextNodePtr *Node) *Node {
	var maxNodePtr *Node = nil
	for nextNodePtr != nil {
		maxNodePtr = nextNodePtr
		nextNodePtr = nextNodePtr.right
	}
	return maxNodePtr
}

type Tree struct {
	root *Node
	size int
}

func New() *Tree {
	return &Tree{
		root: nil,
		size: 0,
	}
}

func (t *Tree) Insert(data Data) {
	node := newNode(data)
	nextNodePtr := t.root
	var parentNodePtr *Node = nil
	for nextNodePtr != nil {
		parentNodePtr = nextNodePtr
		if data.LessThan(nextNodePtr.data) {
			nextNodePtr = nextNodePtr.left
		} else {
			nextNodePtr = nextNodePtr.right
		}
	}
	node.parent = parentNodePtr
	if parentNodePtr == nil {
		t.root = node
	} else {
		if data.LessThan(parentNodePtr.data) {
			parentNodePtr.left = node
		} else {
			parentNodePtr.right = node
		}
	}
	t.size++
}

func (t *Tree) Search(searchData interface{}) *Node {
	nextNodePtr := t.root
	for nextNodePtr != nil {
		if nextNodePtr.data.EqualsTo(searchData) {
			return nextNodePtr
		}
		if nextNodePtr.data.LessThan(searchData) {
			nextNodePtr = nextNodePtr.right
		} else {
			nextNodePtr = nextNodePtr.left
		}
	}
	return nil
}

func (t *Tree) Delete(nodePtr *Node) {
	if nodePtr != nil {
		if nodePtr.left == nil {
			transplant(t, nodePtr, nodePtr.right)
		} else if nodePtr.right == nil {
			transplant(t, nodePtr, nodePtr.left)
		} else {
			x := minimum(nodePtr.right)
			if x.parent != nodePtr {
				transplant(t, x, x.right)
				x.right = nodePtr.right
				x.right.parent = x
			}
			transplant(t, nodePtr, x)
			x.left = nodePtr.left
			x.left.parent = x
		}
		nodePtr.parent, nodePtr.left, nodePtr.right = nil, nil, nil
		t.size--
	}
}

func transplant(t *Tree, u, v *Node) {
	if u.parent == nil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	if v != nil {
		v.parent = u.parent
	}
}

func (t *Tree) Minimum() *Node {
	return minimum(t.root)
}

func (t *Tree) Maximum() *Node {
	return maximum(t.root)
}

type sortedData struct {
	data       []Data
	currentIdx int
}

func (t *Tree) InorderTreeWalk() []Data {
	sortedData := sortedData{
		data:       make([]Data, t.size),
		currentIdx: 0,
	}
	inorderTreeWalk(t.root, &sortedData)
	return sortedData.data
}

func inorderTreeWalk(nodePtr *Node, sortedData *sortedData) {
	if nodePtr != nil {
		inorderTreeWalk(nodePtr.left, sortedData)
		sortedData.data[sortedData.currentIdx] = nodePtr.data
		sortedData.currentIdx++
		inorderTreeWalk(nodePtr.right, sortedData)
	}
}
