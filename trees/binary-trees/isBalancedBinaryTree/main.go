package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func isBalanced(root *TreeNode) bool {

	if dfsHeight(root) != -1 {
		return true
	}
	return false

}

func dfsHeight(root *TreeNode) int {

	if root == nil {
		return 0
	}

	lh := dfsHeight(root.Left)

	if lh == -1 {
		return -1
	}
	rh := dfsHeight(root.Right)

	if rh == -1 {
		return -1
	}

	if abs(lh, rh) > 1 {
		return -1
	}

	return 1 + max(lh, rh)

}

func abs(a, b int) int {
	if a-b < 0 {
		return -(a - b)
	}

	return a - b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}




func main() {

	root := &TreeNode{Val: 3}

	root.Left = &TreeNode{Val: 9}

	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	//root.Left.Right.Left = &TreeNode{Val: 1}

	//fmt.Println(postorderTraversal(root))
	//fmt.Println(levelOrder(root))
	fmt.Println(isBalanced(root))

}

