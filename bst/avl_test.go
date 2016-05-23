package bst

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func kvString(k, v interface{}) string {
	a := k.(int)
	return strconv.Itoa(a)
}

func TestAvlInsert(t *testing.T) {
	tree := NewAvl(intCompare)

	// right rotate
	//      22        21
	//     /         / \
	//    21        20 22
	//    /
	//   20
	tree.Clear()
	tree.Insert(22, 1)
	tree.Insert(21, 1)
	tree.Insert(20, 1)
	if !checkNode(tree.root, 21, 20, 22) {
		t.Errorf("insert")
	}

	// left rotate
	//      20          21
	//       \         / \
	//       21       20 22
	//        \
	//        22
	tree.Clear()
	tree.Insert(20, 1)
	tree.Insert(21, 1)
	tree.Insert(22, 1)
	if !checkNode(tree.root, 21, 20, 22) {
		t.Errorf("insert")
	}

	// left rotate left subtree
	// right rotate
	//       9        9     8
	//      /        /     / \
	//     7        8     7   9
	//      \      /
	//       8    7
	tree.Clear()
	tree.Insert(9, 1)
	tree.Insert(7, 1)
	tree.Insert(8, 1)
	if !checkNode(tree.root, 8, 7, 9) {
		t.Errorf("insert")
	}

	// right rotate right subtree
	// left rotate
	//       7      7         8
	//        \      \       / \
	//         9      8     7   9
	//        /        \
	//       8          9
	tree.Clear()
	tree.Insert(7, 1)
	tree.Insert(9, 1)
	tree.Insert(8, 1)
	if !checkNode(tree.root, 8, 7, 9) {
		t.Errorf("insert")
	}

	// right rotate
	//         25
	//        /  \
	//       20   30
	//      / \   / \
	//     16 22 29 31
	//     /\
	//    12 18
	//    /
	//   10
	//
	//                    25
	//                   /  \
	//                 16    30
	//                 /\    /\
	//               12 20 29 31
	//               /  /\
	//              10 18 22
	tree.Clear()
	tree.Insert(25, 1)
	tree.Insert(20, 1)
	tree.Insert(30, 1)
	tree.Insert(16, 1)
	tree.Insert(22, 1)
	tree.Insert(29, 1)
	tree.Insert(31, 1)
	tree.Insert(12, 1)
	tree.Insert(18, 1)
	tree.Insert(10, 1)
	if !checkNode(tree.root, 25, 16, 30) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left, 16, 12, 20) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right, 30, 29, 31) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left.left, 12, 10, 0) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left.right, 20, 18, 22) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right.left, 29, 0, 0) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right.right, 31, 0, 0) {
		t.Errorf("insert")
	}

	// left rotate
	//           25
	//          /  \
	//        20    30
	//        /\    /\
	//      16 22 29 31
	//         /\
	//       21 24
	//          /
	//         23
	//
	//                  25
	//                 /  \
	//               22    30
	//              /  \   / \
	//             20  24 29 31
	//            / \  /
	//           16 21 23
	tree.Clear()
	tree.Insert(25, 1)
	tree.Insert(20, 1)
	tree.Insert(30, 1)
	tree.Insert(16, 1)
	tree.Insert(22, 1)
	tree.Insert(29, 1)
	tree.Insert(31, 1)
	tree.Insert(21, 1)
	tree.Insert(24, 1)
	tree.Insert(23, 1)
	if !checkNode(tree.root, 25, 22, 30) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left, 22, 20, 24) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right, 30, 29, 31) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left.left, 20, 16, 21) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left.right, 24, 23, 0) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right.left, 29, 0, 0) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right.right, 31, 0, 0) {
		t.Errorf("insert")
	}

	// left rotate left subtree
	// right rotate
	//             30
	//            /  \
	//          14    35
	//          /\    /\
	//        12 20 33 37
	//       /  / \
	//      11 18 22
	//            /
	//           21
	//
	//
	//               20
	//              /  \
	//            14    30
	//            / \   / \
	//           12 18 22 35
	//           /     /  /\
	//          11    21 33 37
	tree.Clear()
	tree.Insert(30, 1)
	tree.Insert(14, 1)
	tree.Insert(35, 1)
	tree.Insert(12, 1)
	tree.Insert(20, 1)
	tree.Insert(33, 1)
	tree.Insert(37, 1)
	tree.Insert(11, 1)
	tree.Insert(18, 1)
	tree.Insert(22, 1)
	tree.Insert(21, 1)
	if !checkNode(tree.root, 20, 14, 30) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left, 14, 12, 18) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right, 30, 22, 35) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left.left, 12, 11, 0) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left.right, 18, 0, 0) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right.left, 22, 21, 0) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right.right, 35, 33, 37) {
		t.Errorf("insert")
	}

	//            30
	//           /  \
	//         14    40
	//         / \   / \
	//        12 20 35 45
	//            /\  \
	//          33 37 48
	//              \
	//              38
	//
	//                        35
	//                       /  \
	//                     30    40
	//                    / \    / \
	//                   14 33  37  45
	//                   /\      \   \
	//                  12 20    38  48
	tree.Clear()
	tree.Insert(30, 1)
	tree.Insert(14, 1)
	tree.Insert(40, 1)
	tree.Insert(12, 1)
	tree.Insert(20, 1)
	tree.Insert(35, 1)
	tree.Insert(45, 1)
	tree.Insert(33, 1)
	tree.Insert(37, 1)
	tree.Insert(48, 1)
	tree.Insert(38, 1)
	if !checkNode(tree.root, 35, 30, 40) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left, 30, 14, 33) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right, 40, 37, 45) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left.left, 14, 12, 20) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.left.right, 33, 0, 0) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right.left, 37, 0, 38) {
		t.Errorf("insert")
	}
	if !checkNode(tree.root.right.right, 45, 0, 48) {
		t.Errorf("insert")
	}
}

func TestAvlDel(t *testing.T) {
	tree := NewAvl(intCompare)

	//       8          8           9
	//      / \          \         / \
	//     7   9          9       8   10
	//          \          \
	//          10         10
	tree.Clear()
	tree.Insert(8, 1)
	tree.Insert(7, 1)
	tree.Insert(9, 1)
	tree.Insert(10, 1)
	tree.Del(7)
	if !checkNode(tree.root, 9, 8, 10) {
		t.Errorf("del %v %v %v", tree.root, tree.root.right, tree.root.right.right)
	}

	//       8          8           7
	//      / \        /           / \
	//     7   9      7           6   8
	//    /          /
	//   6          6
	tree.Clear()
	tree.Insert(8, 1)
	tree.Insert(7, 1)
	tree.Insert(9, 1)
	tree.Insert(6, 1)
	tree.Del(9)
	if !checkNode(tree.root, 7, 6, 8) {
		t.Errorf("del")
	}

	//       8          8          10
	//      / \          \         / \
	//     7   12        12       8   12
	//         /         /
	//        10        10
	tree.Clear()
	tree.Insert(8, 1)
	tree.Insert(7, 1)
	tree.Insert(12, 1)
	tree.Insert(10, 1)
	tree.Del(7)
	if !checkNode(tree.root, 10, 8, 12) {
		t.Errorf("del")
	}

	//       8          8           6
	//      / \        /           / \
	//     5   9      5           5   8
	//      \          \
	//       6          6
	tree.Clear()
	tree.Insert(8, 1)
	tree.Insert(5, 1)
	tree.Insert(9, 1)
	tree.Insert(6, 1)
	tree.Del(9)
	if !checkNode(tree.root, 6, 5, 8) {
		t.Errorf("del")
	}

	//       8          8              15
	//      / \        / \            / \
	//     7  15      6   15         8   16
	//    /   / \         / \       / \   \
	//   6   10 16       10 16     6  10  17
	//           \           \
	//           17          17
	tree.Clear()
	tree.Insert(8, 1)
	tree.Insert(7, 1)
	tree.Insert(15, 1)
	tree.Insert(6, 1)
	tree.Insert(10, 1)
	tree.Insert(16, 1)
	tree.Insert(17, 1)
	tree.Del(7)
	if !checkNode(tree.root, 15, 8, 16) {
		t.Errorf("del %v", *(tree.root))
	}
	if !checkNode(tree.root.left, 8, 6, 10) {
		t.Errorf("del %v", *(tree.root.left))
	}
	if !checkNode(tree.root.right, 16, 0, 17) {
		t.Errorf("del %v", *(tree.root.right))
	}

	//       10                  10              6
	//      /  \                /  \            / \
	//     6    12             6   13          5   10
	//    / \    \            / \             /    / \
	//   5   7    13         5   7           3    7  13
	//  /                   /
	// 3                   3
	tree.Clear()
	tree.Insert(10, 1)
	tree.Insert(6, 1)
	tree.Insert(12, 1)
	tree.Insert(5, 1)
	tree.Insert(7, 1)
	tree.Insert(13, 1)
	tree.Insert(3, 1)
	tree.Del(12)
	if !checkNode(tree.root, 6, 5, 10) {
		t.Errorf("del %v", *(tree.root))
	}
	if !checkNode(tree.root.left, 5, 3, 0) {
		t.Errorf("del %v", *(tree.root.left))
	}
	if !checkNode(tree.root.right, 10, 7, 13) {
		t.Errorf("del %v", *(tree.root.right))
	}

	//       10                  10                10                 7
	//      /  \                /  \              /  \               / \
	//     6    12             6   13            7   13             6  10
	//    / \    \            / \               / \                /   / \
	//   5   7    13         5   7             6   8              5   8  13
	//        \                   \           /
	//         8                   8         5
	tree.Clear()
	tree.Insert(10, 1)
	tree.Insert(6, 1)
	tree.Insert(12, 1)
	tree.Insert(5, 1)
	tree.Insert(7, 1)
	tree.Insert(13, 1)
	tree.Insert(8, 1)
	tree.Del(12)
	if !checkNode(tree.root, 7, 6, 10) {
		t.Errorf("del %v", *(tree.root))
	}
	if !checkNode(tree.root.left, 6, 5, 0) {
		t.Errorf("del %v", *(tree.root.left))
	}
	if !checkNode(tree.root.right, 10, 8, 13) {
		t.Errorf("del %v", *(tree.root.right))
	}

	//       8          8              8                   10
	//      / \        / \            / \                 /  \
	//     7  15      6   15         6   10              8   15
	//    /   / \         / \            / \            / \   \
	//   6   10 17       10 17          9  15          6   9  17
	//       /           /                   \
	//       9          9                    17
	tree.Clear()
	tree.Insert(8, 1)
	tree.Insert(7, 1)
	tree.Insert(15, 1)
	tree.Insert(6, 1)
	tree.Insert(10, 1)
	tree.Insert(17, 1)
	tree.Insert(9, 1)
	tree.Del(7)
	if !checkNode(tree.root, 10, 8, 15) {
		t.Errorf("del %v", *(tree.root))
	}
	if !checkNode(tree.root.left, 8, 6, 9) {
		t.Errorf("del %v", *(tree.root.left))
	}
	if !checkNode(tree.root.right, 15, 0, 17) {
		t.Errorf("del %v", *(tree.root.right))
	}
}

func TestAvl(t *testing.T) {
	size := 10000
	keys := make([]int, 0, size)

	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < size; i++ {
		k := r.Int() % (size * 2)
		keys = append(keys, k)
	}

	tree := NewAvl(intCompare)
	for i := 0; i < size; i++ {
		tree.Insert(keys[i], 1)
	}

	if !avlCheckHeight(tree.root) {
		t.Error("height")
	}

	for i := 0; i < size*10; i++ {
		key := keys[r.Int()%size]
		if _, found := tree.Find(key); !found {
			t.Error("find")
		}
	}

	for i := 0; i < size/2; i++ {
		key := keys[r.Int()%size]
		tree.Del(key)
	}

	if !avlCheckHeight(tree.root) {
		t.Error("height")
		return
	}

}

func avlCheckHeight(root *node) bool {
	if root == nil {
		return true
	}

	h := root.left.height() - root.right.height()
	if h >= 2 || h <= -2 {
		return false
	}

	if !avlCheckHeight(root.left) {
		return false
	}

	if !avlCheckHeight(root.right) {
		return false
	}

	return true
}
