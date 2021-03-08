package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	v1, v2 := dfsRobTree(root)

	return max(v1, v2)
}

func dfsRobTree(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	l0, l1 := dfsRobTree(root.Left)
	r0, r1 := dfsRobTree(root.Right)

	tmp0 := max(l0, l1) + max(r0, r1) // do not rob current node
	tmp1 := root.Val + l0 + r0        // rob current node

	return tmp0, tmp1
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func main() {
	// res := rob([]int{1, 2, 3, 1})
	root := TreeNode{Val: 3, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 1}}}
	res := rob(&root)
	fmt.Println(res)
}
