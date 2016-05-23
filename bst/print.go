package bst

import (
	"bytes"
	"container/list"
	"math"
)

func printChars(c byte, width int, output *bytes.Buffer) {
	for i := 0; i < width; i++ {
		output.WriteByte(c)
	}
}

func printBranches(branchLen int, nodeSpaceLen int, startLen int, nodesInThisLevel int, nodeList *list.List, output *bytes.Buffer) {
	e := nodeList.Front()
	for i := 0; i < nodesInThisLevel/2; i++ {
		if i == 0 {
			printChars(' ', startLen-1, output)
		} else {
			printChars(' ', nodeSpaceLen-2, output)
		}

		if e.Value != nil && e.Value.(*node) != nil {
			output.WriteByte('/')
		} else {
			output.WriteByte(' ')
		}

		e = e.Next()

		printChars(' ', 2*branchLen+2, output)
		if e.Value != nil && e.Value.(*node) != nil {
			output.WriteByte('\\')
		} else {
			output.WriteByte(' ')
		}
		e = e.Next()
	}
	output.WriteByte('\n')
}

func printNodes(stringFunc func(interface{}, interface{}, int) string,
	branchLen int, nodeSpaceLen int, startLen int,
	nodesInThisLevel int, nodeList *list.List, output *bytes.Buffer) {
	e := nodeList.Front()
	for i := 0; i < nodesInThisLevel; i++ {
		if i == 0 {
			printChars(' ', startLen, output)
		} else {
			printChars(' ', nodeSpaceLen, output)
		}

		var n *node
		if e.Value != nil {
			n = e.Value.(*node)
		}

		var fill byte = ' '
		if n != nil && n.left != nil {
			fill = '_'
		}

		width := branchLen + 2
		if n != nil {
			str := stringFunc(n.key, n.val, n.addition)
			printChars(fill, width-len(str), output)
			output.WriteString(str)
		} else {
			printChars(fill, width, output)
		}

		if n != nil && n.right != nil {
			printChars('_', branchLen, output)
		} else {
			printChars(' ', branchLen, output)
		}

		e = e.Next()
	}
	output.WriteByte('\n')
}

func printLeaves(stringFunc func(interface{}, interface{}, int) string,
	indentSpace int, level int, nodesInThisLevel int,
	nodeList *list.List, output *bytes.Buffer) {
	e := nodeList.Front()
	for i := 0; i < nodesInThisLevel; i++ {
		width := 2*level + 2
		if i == 0 {
			width = indentSpace + 2
		}

		var n *node
		if e.Value != nil {
			n = e.Value.(*node)
		}

		if n != nil {
			str := stringFunc(n.key, n.val, n.addition)
			printChars(' ', width-len(str), output)
			output.WriteString(str)
		} else {
			printChars(' ', width, output)
		}

		e = e.Next()
	}
	output.WriteByte('\n')
}

func calcTreeHeight(root *node) int {
	if root == nil {
		return 0
	}

	leftHeight := calcTreeHeight(root.left)
	rifhtHeight := calcTreeHeight(root.right)

	if leftHeight > rifhtHeight {
		return leftHeight + 1
	} else {
		return rifhtHeight + 1
	}
}

func (tree *UnbalancedTree) String(level int, stringFunc func(interface{}, interface{}, int) string) string {
	if tree.root == nil {
		return ""
	}

	indentSpace := 1

	h := calcTreeHeight(tree.root)
	nodesInThisLevel := 1
	tmp := int(math.Pow(2.0, float64(h)))

	// eq of the length of branch for each node of each level
	branchLen := 2*(tmp-1) - (3-level)*int(math.Pow(2.0, float64(h-1)))
	// distance between left neighbor node's right arm and right neighbor node's left arm
	nodeSpaceLen := 2 + (level+1)*tmp
	// starting space to the first node to print of each level (for the left most node of each level only)
	startLen := branchLen + (3 - level) + indentSpace

	nodeList := list.New()
	nodeList.PushBack(tree.root)

	var output bytes.Buffer
	for r := 1; r < h; r++ {
		printBranches(branchLen, nodeSpaceLen, startLen, nodesInThisLevel, nodeList, &output)
		branchLen = branchLen/2 - 1
		nodeSpaceLen = nodeSpaceLen/2 + 1
		startLen = branchLen + (3 - level) + indentSpace
		printNodes(stringFunc, branchLen, nodeSpaceLen, startLen, nodesInThisLevel, nodeList, &output)

		for i := 0; i < nodesInThisLevel; i++ {
			frontElem := nodeList.Front()
			nodeList.Remove(frontElem)
			var currNode *node
			if frontElem.Value != nil {
				currNode = frontElem.Value.(*node)
			}

			if currNode != nil {
				nodeList.PushBack(currNode.left)
				nodeList.PushBack(currNode.right)
			} else {
				nodeList.PushBack(nil)
				nodeList.PushBack(nil)
			}
		}
		nodesInThisLevel *= 2
	}
	printBranches(branchLen, nodeSpaceLen, startLen, nodesInThisLevel, nodeList, &output)
	printLeaves(stringFunc, indentSpace, level, nodesInThisLevel, nodeList, &output)

	return output.String()
}
