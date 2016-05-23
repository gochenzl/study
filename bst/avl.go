package bst

type AvlTree struct {
	UnbalancedTree
}

func NewAvl(comparer Comparer) AvlTree {
	tree := AvlTree{}
	tree.cmp = comparer

	return tree
}

func (tree *AvlTree) Insert(k, v interface{}) {
	p, exist := insert(tree.root, k, v, tree.cmp)
	if !exist {
		tree.size++
	}
	if tree.root == nil {
		tree.root = p
		return
	}

	if exist {
		return
	}

	p.setHeight(1)
	updateHeight(p.parent)

	if p.parent != nil &&
		p.parent.parent != nil {
		p = p.parent.parent
		for p != nil {
			p = rebalance(&(tree.root), p)
		}
	}
}

func (tree *AvlTree) Del(key interface{}) bool {
	p, _, y, _ := del(&(tree.root), key, tree.cmp)
	if y == nil {
		return false
	}

	tree.size--

	if p == nil {
		p = tree.root
	}

	updateHeight(p)

	for p != nil {
		p = rebalance(&(tree.root), p)
	}

	return true
}

//   LL(left-left) right rotate
//
//       A             B
//      /             / \
//     B             C   A
//    /
//   C
//
//   LR(left-right) (1)left rotate to left subtree (2) right rotate
//
//       A             A           C
//      /             /           / \
//     B             C           B   A
//      \           /
//       C         B
//
//   RR(right-right) left rotate
//
//    A             B
//     \           / \
//      B         A   C
//       \
//        C
//
//   RR(right-left) (1)right rotate to right subtree (2) left rotate
//
//    A            A             C
//     \            \           / \
//      B            C         A   B
//     /              \
//    C                B
//
func rebalance(proot **node, p *node) *node {
	if p == nil {
		return nil
	}

	var doRightRotate, doLeftRotate bool

	parent := p.parent

	leftHeight := p.left.height()
	rightHeight := p.right.height()

	if leftHeight-rightHeight == 2 {
		// LL
		if p.left.left.height() >= p.left.right.height() {
			rightRotate(p, proot, true)
		} else { // LR
			p = p.left
			parent = p.parent
			leftRotate(p, proot, true)
			doRightRotate = true
		}

	} else if rightHeight-leftHeight == 2 {
		// RR
		if p.right.right.height() >= p.right.left.height() {
			leftRotate(p, proot, true)
		} else { // RL
			p = p.right
			parent = p.parent
			rightRotate(p, proot, true)
			doLeftRotate = true
		}
	} else {
		return parent
	}

	if doRightRotate || doLeftRotate {
		p = parent
		parent = p.parent
		if doRightRotate {
			rightRotate(p, proot, true)
		} else if doLeftRotate {
			leftRotate(p, proot, true)
		}
	}

	return parent
}
