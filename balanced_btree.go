package practice

type AVLNode struct {
	value  int
	left   *AVLNode
	right  *AVLNode
	bf     int
	height int
}

type AVL struct {
	size int
	root *AVLNode
}

func NewAVL() *AVL {
	t := AVL{0, nil}

	return &t
}
func NewAVLNode(value int) *AVLNode {
	n := AVLNode{
		value, // value  int
		nil,   // left   *AVLNode
		nil,   // right  *AVLNode
		0,     // bf     int
		1,     // height int
	}

	return &n
}

func (t *AVL) IsEmpty() bool {
	return t.size == 0
}
func (t *AVL) Size() int {
	return t.size
}
func (t *AVL) Has(value int) bool {
	return t.find(t.root, value) != nil
}
func (t *AVL) Insert(value int) bool {
	t.root = t.add(t.root, value)
	t.size++
	return true
}
func (t *AVL) Remove(value int) bool {
	if t.IsEmpty() {
		return false
	}

	var ok bool
	t.root, ok = t.remove(t.root, value)
	t.size--
	return ok
}

func (t *AVL) find(node *AVLNode, value int) *AVLNode {
	if node == nil {
		return node
	}

	if value < node.value {
		return t.find(node.left, value)
	}
	if value > node.value {
		return t.find(node.right, value)
	}

	return node
}
func (t *AVL) add(node *AVLNode, value int) *AVLNode {
	if node == nil {
		return NewAVLNode(value)
	}
	if value < node.value {
		node.left = t.add(node.left, value)
	} else { // assert(value > parent.value)
		node.right = t.add(node.right, value)
	}
	t.update(node)
	return t.rebalance(node)
}
func (t *AVL) remove(node *AVLNode, value int) (*AVLNode, bool) {
	if node == nil {
		return nil, false
	}

	var ok bool
	if value < node.value {
		node.left, ok = t.remove(node.left, value)
	} else if value > node.value {
		node.right, ok = t.remove(node.right, value)
	} else {
		ok = true
		if node.right == nil {
			return node.left, ok
		} else if node.left == nil {
			return node.right, ok
		} else {
			if node.right.height > node.left.height {
				move := t.findMin(node.right)
				node.value = move.value
				node.right, _ = t.remove(node.right, move.value)
			} else {
				move := t.findMax(node.left)
				node.value = move.value
				node.left, _ = t.remove(node.left, move.value)
			}
		}
	}

	t.update(node)
	return t.rebalance(node), ok
}
func (t *AVL) findMax(node *AVLNode) *AVLNode {
	for node.right != nil {
		node = node.right
	}
	return node
}
func (t *AVL) findMin(node *AVLNode) *AVLNode {
	for node.left != nil {
		node = node.left
	}
	return node
}
func (t *AVL) update(node *AVLNode) {
	if node == nil {
		return
	}

	node.height = 1
	leftHeight := 0
	if node.left != nil {
		leftHeight = node.left.height
	}
	rightHeight := 0
	if node.right != nil {
		rightHeight = node.right.height
	}
	if leftHeight >= rightHeight {
		node.height += leftHeight
	} else {
		node.height += rightHeight
	}

	node.bf = rightHeight - leftHeight
}
func (t *AVL) rebalance(node *AVLNode) *AVLNode {
	if node == nil {
		return node
	}

	if node.bf == -2 {
		if node.left.bf <= 0 {
			return t.leftLeftCase(node)
		} else {
			return t.leftRightCase(node)
		}
	} else if node.bf == 2 {
		if node.right.bf >= 0 {
			return t.rightRightCase(node)
		} else {
			return t.rightLeftCase(node)
		}
	}

	return node
}
func (t *AVL) rotateRight(node *AVLNode) *AVLNode {
	parent := node.left
	node.left = parent.right
	parent.right = node
	t.update(node)
	t.update(parent)
	return parent
}
func (t *AVL) rotateLeft(node *AVLNode) *AVLNode {
	parent := node.right
	node.right = parent.left
	parent.left = node
	t.update(node)
	t.update(parent)
	return parent
}
func (t *AVL) leftLeftCase(node *AVLNode) *AVLNode {
	return t.rotateRight(node)
}
func (t *AVL) leftRightCase(node *AVLNode) *AVLNode {
	node.left = t.rotateLeft(node.left)
	return t.leftLeftCase(node)
}
func (t *AVL) rightRightCase(node *AVLNode) *AVLNode {
	return t.rotateLeft(node)
}
func (t *AVL) rightLeftCase(node *AVLNode) *AVLNode {
	node.right = t.rotateRight(node.right)
	return t.rightRightCase(node)
}
