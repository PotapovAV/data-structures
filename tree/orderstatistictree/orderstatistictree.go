package orderstatistictree

type Data interface {
	LessThan(other interface{}) bool
	EqualsTo(other interface{}) bool
}

type Node struct {
	data                       Data
	parent, left, right        *Node
	height, subtreeNodesAmount int
}

func newNode(data Data) *Node {
	return &Node{
		data:               data,
		parent:             nil,
		left:               nil,
		right:              nil,
		height:             0,
		subtreeNodesAmount: 1,
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
	t.root = insert(t, t.root, data)
	t.size++
}

func insert(t *Tree, node *Node, data Data) *Node {
	if node == nil {
		return newNode(data)
	}
	if data.LessThan(node.data) {
		node.left = insert(t, node.left, data)
		node.left.parent = node
	} else {
		node.right = insert(t, node.right, data)
		node.right.parent = node
	}

	balanceFactor := getBalanceFactor(node)

	if (balanceFactor == 2) && (data.LessThan(node.left.data)) {
		return rightRotate(t, node)
	}
	if (balanceFactor == -2) && (!data.LessThan(node.right.data)) {
		return leftRotate(t, node)
	}
	if (balanceFactor == 2) && (!data.LessThan(node.left.data)) {
		leftRotate(t, node.left)
		return rightRotate(t, node)
	}
	if (balanceFactor == -2) && (data.LessThan(node.right.data)) {
		rightRotate(t, node.right)
		return leftRotate(t, node)
	}
	node.height = calcNodeHeight(node)
	node.subtreeNodesAmount = calcSubtreeNodesAmount(node)
	return node
}

func getBalanceFactor(n *Node) int {
	if n == nil {
		return 0
	}
	return getHeight(n.left) - getHeight(n.right)
}

func getHeight(n *Node) int {
	if n == nil {
		return -1
	}
	return n.height
}

func calcNodeHeight(n *Node) int {
	leftHeight := getHeight(n.left)
	rightHeight := getHeight(n.right)
	return 1 + max(leftHeight, rightHeight)
}

func calcSubtreeNodesAmount(n *Node) int {
	leftNodesAmount := getNodesAmount(n.left)
	rightNodesAmount := getNodesAmount(n.right)
	return 1 + leftNodesAmount + rightNodesAmount
}

func getNodesAmount(n *Node) int {
	if n == nil {
		return 0
	}
	return n.subtreeNodesAmount
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func leftRotate(t *Tree, x *Node) *Node {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		t.root = y
	} else if x.parent.left == x {
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x
	x.parent = y

	x.height = calcNodeHeight(x)
	x.subtreeNodesAmount = calcSubtreeNodesAmount(x)
	y.height = calcNodeHeight(y)
	y.subtreeNodesAmount = calcSubtreeNodesAmount(y)
	return y
}

func rightRotate(t *Tree, y *Node) *Node {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y.parent.left == y {
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y
	y.parent = x

	y.height = calcNodeHeight(y)
	y.subtreeNodesAmount = calcSubtreeNodesAmount(y)
	x.height = calcNodeHeight(x)
	x.subtreeNodesAmount = calcSubtreeNodesAmount(x)
	return x
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

func (t *Tree) Select(i int) *Node {
	// [0..n-1] -> [1..n]
	return osSelect(t.root, i+1)
}

func osSelect(n *Node, j int) *Node {
	if n == nil {
		return nil
	}
	r := getNodesAmount(n.left) + 1
	if r == j {
		return n
	} else if j < r {
		return osSelect(n.left, j)
	} else {
		return osSelect(n.right, j-r)
	}
}

func (t *Tree) Delete(node *Node) {
	var fixupStartNode *Node = nil
	if node != nil {
		if node.left == nil {
			fixupStartNode = transplant(t, node, node.right)
		} else if node.right == nil {
			fixupStartNode = transplant(t, node, node.left)
		} else {
			x := minimum(node.right)
			if x.parent != node {
				fixupStartNode = transplant(t, x, x.right)
				x.right = node.right
				x.right.parent = x
				deleteFixup(t, fixupStartNode, x)
			}
			transplant(t, node, x)
			x.left = node.left
			x.left.parent = x
			deleteFixup(t, x, nil)
		}
		node.parent, node.left, node.right = nil, nil, nil
		deleteFixup(t, fixupStartNode, nil)
		t.size--
	}
}

func deleteFixup(t *Tree, start *Node, stop *Node) {
	for start != stop {
		start.height = calcNodeHeight(start)
		start.subtreeNodesAmount = calcSubtreeNodesAmount(start)
		balanceFactor := getBalanceFactor(start)
		if balanceFactor == 2 {
			leftSubtreeBalanceFactor := getBalanceFactor(start.left)
			if leftSubtreeBalanceFactor == 1 {
				start = rightRotate(t, start)
			} else {
				leftRotate(t, start.left)
				start = rightRotate(t, start)
			}
		} else if balanceFactor == -2 {
			rightSubtreeBalanceFactor := getBalanceFactor(start.right)
			if rightSubtreeBalanceFactor == -1 {
				start = leftRotate(t, start)
			} else {
				rightRotate(t, start.right)
				start = leftRotate(t, start)
			}
		}
		start = start.parent
	}
}

func transplant(t *Tree, u, v *Node) *Node {
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
	return u.parent
}

func (t *Tree) Minimum() *Node {
	return minimum(t.root)
}

func minimum(nextNodePtr *Node) *Node {
	var minNodePtr *Node = nil
	for nextNodePtr != nil {
		minNodePtr = nextNodePtr
		nextNodePtr = nextNodePtr.left
	}
	return minNodePtr
}

func (t *Tree) Maximum() *Node {
	return maximum(t.root)
}

func maximum(nextNodePtr *Node) *Node {
	var maxNodePtr *Node = nil
	for nextNodePtr != nil {
		maxNodePtr = nextNodePtr
		nextNodePtr = nextNodePtr.right
	}
	return maxNodePtr
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
