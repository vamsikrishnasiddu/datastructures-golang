package main


import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	lh := maxDepth(root.Left)

	rh := maxDepth(root.Right)

	return 1 + max(lh, rh)

}

func max(a,b int)int{

  if a>b{
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
	fmt.Println(maxDepth(root))

}


