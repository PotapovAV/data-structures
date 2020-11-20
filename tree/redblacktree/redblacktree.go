package redblacktree

const (
	red   = iota
	black = iota
)

type Data interface {
	LessThan(other interface{}) bool
	EqualsTo(other interface{}) bool
}

type Tree struct {
	root, tNil *Node
	size       int
}

type Node struct {
	data                Data
	parent, right, left *Node
	color               int
}

func newNode(data Data) *Node {
	return &Node{
		data:   data,
		parent: nil,
		right:  nil,
		left:   nil,
		color:  red,
	}
}

func (n *Node) Data() Data {
	return n.data
}

func New() *Tree {
	tNilNode := newNode(nil)
	tNilNode.color = black
	return &Tree{
		root: tNilNode,
		tNil: tNilNode,
		size: 0,
	}
}

func (t *Tree) GetTNil() *Node {
	return t.tNil
}

func (t *Tree) Insert(data Data) {
	node := newNode(data)
	y := t.tNil
	x := t.root
	for x != t.tNil {
		y = x
		if node.data.LessThan(x.data) {
			x = x.left
		} else {
			x = x.right
		}
	}
	node.parent = y
	if y == t.tNil {
		t.root = node
	} else if node.data.LessThan(y.data) {
		y.left = node
	} else {
		y.right = node
	}
	node.right = t.tNil
	node.left = t.tNil
	t.size++
	t.insertFixup(node)
}

func (t *Tree) insertFixup(node *Node) {
	for node.parent.color == red {
		if node.parent == node.parent.parent.left {
			y := node.parent.parent.right
			if y.color == red {
				node.parent.color = black
				y.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					leftRotate(t, node)
				}
				node.parent.color = black
				node.parent.parent.color = red
				rightRotate(t, node.parent.parent)
			}
		} else {
			y := node.parent.parent.left
			if y.color == red {
				node.parent.color = black
				y.color = black
				node.parent.parent.color = red
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					rightRotate(t, node)
				}
				node.parent.color = black
				node.parent.parent.color = red
				leftRotate(t, node.parent.parent)
			}
		}
	}
	t.root.color = black
}

func leftRotate(t *Tree, x *Node) {
	y := x.right
	x.right = y.left
	if y.left != t.tNil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == t.tNil {
		t.root = y
	} else if x.parent.left == x {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y
}

func rightRotate(t *Tree, y *Node) {
	x := y.left
	y.left = x.right
	if x.right != t.tNil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == t.tNil {
		t.root = x
	} else if y.parent.left == y {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x
}

func (t *Tree) Delete(node *Node) {
	if node != t.tNil {
		y := node
		yColor := y.color
		x := t.tNil
		if node.left == t.tNil {
			x = node.right
			transplant(t, node, node.right)
		} else if node.right == t.tNil {
			x = node.left
			transplant(t, node, node.left)
		} else {
			y = minimum(t, node.right)
			yColor = y.color
			x = y.right
			if y.parent == node {
				x.parent = y
			} else {
				transplant(t, y, y.right)
				y.right = node.right
				y.right.parent = y
			}
			transplant(t, node, y)
			y.left = node.left
			y.left.parent = y
			y.color = node.color
		}
		t.size--
		if yColor == black {
			deleteFixup(t, x)
		}
	}
}

func transplant(t *Tree, u, v *Node) {
	if u.parent == t.tNil {
		t.root = v
	} else if u == u.parent.left {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func deleteFixup(t *Tree, x *Node) {
	for (x != t.root) && (x.color == black) {
		if x == x.parent.left {
			w := x.parent.right
			if w.color == red {
				w.color = black
				x.parent.color = red
				leftRotate(t, x.parent)
				w = x.parent.right
			}
			if (w.left.color == black) && (w.right.color == black) {
				w.color = red
				x = x.parent
			} else {
				if w.right.color == black {
					w.left.color = black
					w.color = red
					rightRotate(t, w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.right.color = black
				leftRotate(t, x.parent)
				x = t.root
			}
		} else {
			w := x.parent.left
			if w.color == red {
				w.color = black
				x.parent.color = black
				rightRotate(t, x.parent)
				w = x.parent.left
			}
			if (w.left.color) == black && (w.right.color == black) {
				w.color = red
				x = x.parent
			} else {
				if w.left.color == black {
					w.right.color = black
					w.color = red
					leftRotate(t, w)
					w = x.parent.right
				}
				w.color = x.parent.color
				x.parent.color = black
				w.left.color = black
				rightRotate(t, x.parent)
				x = t.root
			}
		}
	}
	x.color = black
}

func (t *Tree) Minimum() *Node {
	return minimum(t, t.root)
}

func (t *Tree) Maximum() *Node {
	return maximum(t, t.root)
}

func minimum(t *Tree, nextNodePtr *Node) *Node {
	var minNodePtr *Node = t.tNil
	for nextNodePtr != t.tNil {
		minNodePtr = nextNodePtr
		nextNodePtr = nextNodePtr.left
	}
	return minNodePtr
}

func maximum(t *Tree, nextNodePtr *Node) *Node {
	var maxNodePtr *Node = t.tNil
	for nextNodePtr != t.tNil {
		maxNodePtr = nextNodePtr
		nextNodePtr = nextNodePtr.right
	}
	return maxNodePtr
}

func (t *Tree) Search(searchData interface{}) *Node {
	nextNodePtr := t.root
	for nextNodePtr != t.tNil {
		if nextNodePtr.data.EqualsTo(searchData) {
			return nextNodePtr
		}
		if nextNodePtr.data.LessThan(searchData) {
			nextNodePtr = nextNodePtr.right
		} else {
			nextNodePtr = nextNodePtr.left
		}
	}
	return t.tNil
}

func (t *Tree) Successor(n *Node) *Node {
	if n.right != t.tNil {
		return minimum(t, n.right)
	}
	parentPtr := n.parent
	for (parentPtr != t.tNil) && (n == parentPtr.right) {
		n = parentPtr
		parentPtr = parentPtr.parent
	}
	return parentPtr
}

func (t *Tree) Predecessor(n *Node) *Node {
	if n.left != t.tNil {
		return maximum(t, n.left)
	}
	parentPtr := n.parent
	for (parentPtr != t.tNil) && (n == parentPtr.left) {
		n = parentPtr
		parentPtr = parentPtr.parent
	}
	return parentPtr
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
	inorderTreeWalk(t, t.root, &sortedData)
	return sortedData.data
}

func inorderTreeWalk(t *Tree, nodePtr *Node, sortedData *sortedData) {
	if nodePtr != t.tNil {
		inorderTreeWalk(t, nodePtr.left, sortedData)
		sortedData.data[sortedData.currentIdx] = nodePtr.data
		sortedData.currentIdx++
		inorderTreeWalk(t, nodePtr.right, sortedData)
	}
}
