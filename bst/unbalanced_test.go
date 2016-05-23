package bst

import (
	"math/rand"
	"testing"
	"time"
)

func intCompare(a, b interface{}) int {
	ia := a.(int)
	ib := b.(int)

	if ia < ib {
		return -1
	}

	if ia > ib {
		return 1
	}

	return 0
}

func inArray(val int, a []int) bool {
	for i := 0; i < len(a); i++ {
		if val == a[i] {
			return true
		}
	}

	return false
}

func checkNode(n *node, key int, lkey int, rkey int) bool {
	if n == nil {
		return false
	}

	if n.key != key {
		return false
	}

	if lkey == 0 && n.left != nil {
		return false
	}

	if rkey == 0 && n.right != nil {
		return false
	}

	if lkey != 0 && n.left.key != lkey {
		return false
	}

	if rkey != 0 && n.right.key != rkey {
		return false
	}

	return true
}

func TestUnbalancedInsert(t *testing.T) {
	tree := &UnbalancedTree{cmp: intCompare}

	//              10
	tree.Insert(10, 10)
	if !checkNode(tree.root, 10, 0, 0) {
		t.Errorf("insert 10")
	}

	//              10
	//             /
	//            8
	tree.Insert(8, 8)
	if !checkNode(tree.root, 10, 8, 0) {
		t.Errorf("insert 8")
	}
	if !checkNode(tree.root.left, 8, 0, 0) {
		t.Errorf("insert 8")
	}

	//              10
	//             /  \
	//            8   11
	tree.Insert(11, 11)
	if !checkNode(tree.root, 10, 8, 11) {
		t.Errorf("insert 8")
	}
	if !checkNode(tree.root.left, 8, 0, 0) {
		t.Errorf("insert 8")
	}
	if !checkNode(tree.root.right, 11, 0, 0) {
		t.Errorf("insert 8")
	}

	//              10
	//             /  \
	//            8   11
	//             \
	//              9
	tree.Insert(9, 9)
	if !checkNode(tree.root, 10, 8, 11) {
		t.Errorf("insert 8")
	}
	if !checkNode(tree.root.left, 8, 0, 9) {
		t.Errorf("insert 8")
	}
	if !checkNode(tree.root.right, 11, 0, 0) {
		t.Errorf("insert 8")
	}
	if !checkNode(tree.root.left.right, 9, 0, 0) {
		t.Errorf("insert 8")
	}

	//              10
	//             /  \
	//            8   11
	//           / \
	//          6   9
	tree.Insert(6, 6)
	if !checkNode(tree.root, 10, 8, 11) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.left, 8, 6, 9) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.right, 11, 0, 0) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.left.right, 9, 0, 0) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.left.left, 6, 0, 0) {
		t.Errorf("insert 6")
	}

	//              10
	//             /  \
	//            8   11
	//           / \    \
	//          6   9    20
	tree.Insert(20, 20)
	if !checkNode(tree.root, 10, 8, 11) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.left, 8, 6, 9) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.right, 11, 0, 20) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.left.right, 9, 0, 0) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.left.left, 6, 0, 0) {
		t.Errorf("insert 6")
	}
	if !checkNode(tree.root.right.right, 20, 0, 0) {
		t.Errorf("insert 6")
	}
}

func TestUnbalancedDel(t *testing.T) {
	tree := &UnbalancedTree{}
	tree.cmp = intCompare

	tree.Insert(1, 1)
	tree.Del(1)
	if tree.root != nil {
		t.Errorf("del a node with no children")
	}

	tree.Clear()
	tree.Insert(2, 2)
	tree.Insert(1, 1)
	tree.Del(1)
	if !checkNode(tree.root, 2, 0, 0) {
		t.Errorf("del a node with no children")
	}

	tree.Clear()
	tree.Insert(2, 2)
	tree.Insert(1, 1)
	tree.Del(2)
	if !checkNode(tree.root, 1, 0, 0) {
		t.Errorf("del a node with left children")
	}

	tree.Clear()
	tree.Insert(1, 1)
	tree.Insert(2, 2)
	tree.Del(1)
	if !checkNode(tree.root, 2, 0, 0) {
		t.Errorf("del a node with right children")
	}

	//    2         3
	//   / \       /
	//   1  3      1
	tree.Clear()
	tree.Insert(2, 2)
	tree.Insert(1, 1)
	tree.Insert(3, 3)
	tree.Del(2)
	if !checkNode(tree.root, 3, 1, 0) ||
		!checkNode(tree.root.left, 1, 0, 0) {
		t.Errorf("del a node with right children")
	}

	//    2         3
	//   / \       / \
	//   1  4      1  4
	//     /
	//    3
	tree.Clear()
	tree.Insert(2, 2)
	tree.Insert(1, 1)
	tree.Insert(4, 4)
	tree.Insert(3, 3)
	tree.Del(2)
	if !checkNode(tree.root, 3, 1, 4) ||
		!checkNode(tree.root.left, 1, 0, 0) ||
		!checkNode(tree.root.right, 4, 0, 0) {
		t.Errorf("del a node with right children")
	}
}

func TestUnbalanced(t *testing.T) {
	tree := &UnbalancedTree{}
	tree.cmp = intCompare

	var keys []int
	kvs := make(map[int]int)

	size := 10000

	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < size; i++ {
		k := rand.Int()
		keys = append(keys, k)
		kvs[k] = rand.Int()
	}

	for k, v := range kvs {
		tree.Insert(k, v)
	}

	for i := 0; i < size/10; i++ {
		key := keys[rand.Int()%size]
		tree.Insert(key, kvs[key])
	}

	for i := 0; i < size*10; i++ {
		key := keys[rand.Int()%size]

		if val, found := tree.Find(key); !found || val.(int) != kvs[key] {
			t.Errorf("Find")
		}
	}

	// find not exist
	for i := 0; i < 1000; i++ {
		var key int
		for {
			key = rand.Int()
			if !inArray(key, keys) {
				break
			}
		}

		if _, found := tree.Find(key); found {
			t.Errorf("Find")
		}
	}

	var delKeys []int
	delCount := 0
	for i := 0; i < size/4; i++ {
		key := keys[rand.Int()%size]
		delKeys = append(delKeys, key)
		if tree.Del(key) {
			delCount++
		}

		if _, found := tree.Find(key); found {
			t.Errorf("Find")
		}
	}

	iter := tree.NewIterator()
	count := 0
	for iter.Next() {
		count++
	}
	if tree.Size() != count {
		t.Error("Del")
	}

	for i := 0; i < 1000; i++ {
		var key int
		for {
			key = keys[rand.Int()%size]
			if !inArray(key, delKeys) {
				break
			}
		}

		if _, found := tree.Find(key); !found {
			t.Errorf("can not find %d", key)
		}
	}
}
